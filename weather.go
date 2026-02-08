package main

import (
	"machine"
	"math"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

// I2C address of the ESP32-S3 weather bridge
const weatherI2CAddr = 0x42

// I2C command codes
const cmdSetLocation = 0x01

// Location represents a weather location
type Location struct {
	Name string
	Lat  float32
	Lon  float32
}

// Available weather locations
//
// To configure your own location:
//  1. Find your city on Wikipedia and look for the coordinates (e.g. 48°51′24″N 2°21′08″E for Paris)
//  2. Convert degrees/minutes/seconds to decimal: degrees + minutes/60 + seconds/3600
//     Example for Paris latitude: 48 + 51/60 + 24/3600 = 48.8567
//     Example for Paris longitude: 2 + 21/60 + 8/3600 = 2.3522
//  3. Use negative values for West longitude or South latitude
var weatherLocations = []Location{
	{"San Francisco", 37.7749, -122.4194},
	{"Paris", 48.8567, 2.3522},
	{"Bengaluru", 12.9789, 77.5917},
	{"Beijing", 39.9067, 116.3975},
	{"Tokyo", 35.6897, 139.6922},
}

// Last selected location index (remembered between menu visits)
var lastSelectedLocation = 0

// WeatherData matches the struct from the ESP32-S3 (32 bytes)
type WeatherData struct {
	CurrentTemp int8
	CurrentCode uint8
	WindSpeed   uint8
	TodayHigh   int8
	TodayLow    int8
	TodayPrecip uint8
	TodayCode   uint8
	Day1High    int8
	Day1Low     int8
	Day1Code    uint8
	Day2High    int8
	Day2Low     int8
	Day2Code    uint8
	Status      uint8 // 0=no data, 1=valid, 2=error
	LastUpdate  uint8 // Minutes since last update
	Reserved    [17]uint8
}

// sendLocationToESP32 sends latitude/longitude to the ESP32-S3 via I2C
// Format: cmd(1) + lat(4, float32 little-endian) + lon(4, float32 little-endian) = 9 bytes
func sendLocationToESP32(loc *Location) error {
	data := make([]byte, 9)
	data[0] = cmdSetLocation

	// Send latitude as float32 (little-endian)
	latBits := math.Float32bits(loc.Lat)
	data[1] = byte(latBits)
	data[2] = byte(latBits >> 8)
	data[3] = byte(latBits >> 16)
	data[4] = byte(latBits >> 24)

	// Send longitude as float32 (little-endian)
	lonBits := math.Float32bits(loc.Lon)
	data[5] = byte(lonBits)
	data[6] = byte(lonBits >> 8)
	data[7] = byte(lonBits >> 16)
	data[8] = byte(lonBits >> 24)

	return machine.I2C0.Tx(weatherI2CAddr, data, nil)
}

// selectWeatherLocation shows a menu to select a weather location
// Returns the selected location index, or -1 if cancelled
func selectWeatherLocation() int {
	display.FillScreen(colors[BLACK])
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 25, "Select Location", colors[YELLOW])
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 230, "A:Select  B:Cancel", colors[TEXT])

	selected := lastSelectedLocation
	prevSelected := -1 // Force initial draw

	for {
		// Only redraw when selection changes
		if selected != prevSelected {
			y := int16(60)
			for i, loc := range weatherLocations {
				// Clear line
				display.FillRectangle(0, y-15, 320, 25, colors[BLACK])

				// Highlight selected
				if i == selected {
					display.FillRectangle(5, y-15, 310, 25, colors[PURPLE])
					tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 15, y, "> "+loc.Name, colors[WHITE])
				} else {
					tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 15, y, "  "+loc.Name, colors[TEXT])
				}
				y += 30
			}
			prevSelected = selected
		}

		// Handle button input
		if !btnUp.Get() {
			selected--
			if selected < 0 {
				selected = len(weatherLocations) - 1
			}
			time.Sleep(200 * time.Millisecond)
		}
		if !btnDown.Get() {
			selected++
			if selected >= len(weatherLocations) {
				selected = 0
			}
			time.Sleep(200 * time.Millisecond)
		}
		if !btnA.Get() {
			time.Sleep(200 * time.Millisecond)
			lastSelectedLocation = selected // Remember for next time
			return selected
		}
		if !btnB.Get() {
			time.Sleep(200 * time.Millisecond)
			return -1
		}

		time.Sleep(50 * time.Millisecond)
	}
}

// Weather displays weather information from the ESP32-S3 bridge
func Weather() {
	for {
		// Show location selection menu
		locIdx := selectWeatherLocation()
		if locIdx < 0 {
			return // User cancelled - exit Weather mode
		}

		loc := &weatherLocations[locIdx]

		// Send location to ESP32
		sendLocationToESP32(loc)

		// Show weather display until user presses B
		if !showWeatherDisplay(loc.Name) {
			return // User wants to exit Weather mode entirely
		}
		// Otherwise loop back to location selection
	}
}

// showWeatherDisplay shows weather for the given location
// Returns true if user pressed B (go back to menu), false if should exit entirely
func showWeatherDisplay(locationName string) bool {
	display.FillScreen(colors[BLACK])

	// Title with location name
	title := "Weather - " + locationName
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 25, title, colors[YELLOW])

	// Instructions
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 230, "B: Change location", colors[TEXT])

	// Show loading message immediately
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 100, "Loading weather data...", colors[TEXT])

	var data WeatherData
	var prevData WeatherData
	lastRead := time.Now().Add(-10 * time.Second) // Force immediate read
	waitingForNewData := true                     // We just sent a new location

	for {
		// Read weather data every 3 seconds
		if time.Since(lastRead) >= 3*time.Second {
			err := readWeatherData(&data)
			lastRead = time.Now()

			// Check if we got fresh data (lastUpdate == 0 means just fetched)
			if waitingForNewData && data.Status == 1 && data.LastUpdate == 0 {
				waitingForNewData = false
			}

			// Only redraw if data changed or we're still waiting
			if waitingForNewData || data != prevData {
				// Clear content area
				display.FillRectangle(0, 35, 320, 180, colors[BLACK])

				if err != nil {
					tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 60, "ESP32 not found", colors[RED])
					tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 85, "Check Qwiic connection", colors[TEXT])
				} else if waitingForNewData || data.Status == 0 {
					tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 100, "Loading weather data...", colors[TEXT])
				} else if data.Status == 2 {
					tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 60, "Weather fetch error", colors[RED])
					tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 85, "Check WiFi connection", colors[TEXT])
				} else {
					drawWeatherDisplay(&data)
				}
				prevData = data
			}
		}

		// Check for back to location menu
		if !btnB.Get() {
			time.Sleep(200 * time.Millisecond) // Debounce
			return true                        // Go back to location selection
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// readWeatherData reads the weather struct from ESP32-S3 via I2C
func readWeatherData(data *WeatherData) error {
	buf := make([]byte, 32)
	err := machine.I2C0.Tx(weatherI2CAddr, nil, buf)
	if err != nil {
		return err
	}

	// Parse the buffer into the struct
	data.CurrentTemp = int8(buf[0])
	data.CurrentCode = buf[1]
	data.WindSpeed = buf[2]
	data.TodayHigh = int8(buf[3])
	data.TodayLow = int8(buf[4])
	data.TodayPrecip = buf[5]
	data.TodayCode = buf[6]
	data.Day1High = int8(buf[7])
	data.Day1Low = int8(buf[8])
	data.Day1Code = buf[9]
	data.Day2High = int8(buf[10])
	data.Day2Low = int8(buf[11])
	data.Day2Code = buf[12]
	data.Status = buf[13]
	data.LastUpdate = buf[14]

	return nil
}

// drawWeatherDisplay renders the weather data on screen
func drawWeatherDisplay(data *WeatherData) {
	y := int16(55)

	// Current conditions
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, y, "Now:", colors[WHITE])
	currentStr := itoa(int(data.CurrentTemp)) + "C  " + weatherCodeToString(data.CurrentCode)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 60, y, currentStr, colors[YELLOW])
	y += 22

	windStr := "Wind: " + itoa(int(data.WindSpeed)) + " km/h"
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 60, y, windStr, colors[TEXT])
	y += 28

	// Today's forecast
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, y, "Today:", colors[WHITE])
	todayStr := itoa(int(data.TodayHigh)) + "/" + itoa(int(data.TodayLow)) + "C"
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 80, y, todayStr, colors[ORANGE])
	y += 22

	precipStr := "Precip: " + itoa(int(data.TodayPrecip)) + "mm"
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 80, y, precipStr, colors[TEXT])
	y += 28

	// Tomorrow
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, y, "Tomorrow:", colors[WHITE])
	day1Str := itoa(int(data.Day1High)) + "/" + itoa(int(data.Day1Low)) + "C " + weatherCodeToEmoji(data.Day1Code)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 110, y, day1Str, colors[TEXT])
	y += 22

	// Day after
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, y, "Day After:", colors[WHITE])
	day2Str := itoa(int(data.Day2High)) + "/" + itoa(int(data.Day2Low)) + "C " + weatherCodeToEmoji(data.Day2Code)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 110, y, day2Str, colors[TEXT])
}

// weatherCodeToString converts WMO weather code to description
func weatherCodeToString(code uint8) string {
	switch code {
	case 0:
		return "Clear"
	case 1, 2, 3:
		return "Partly Cloudy"
	case 45, 48:
		return "Foggy"
	case 51, 53, 55:
		return "Drizzle"
	case 56, 57:
		return "Freezing Drizzle"
	case 61, 63, 65:
		return "Rain"
	case 66, 67:
		return "Freezing Rain"
	case 71, 73, 75:
		return "Snow"
	case 77:
		return "Snow Grains"
	case 80, 81, 82:
		return "Showers"
	case 85, 86:
		return "Snow Showers"
	case 95:
		return "Thunderstorm"
	case 96, 99:
		return "Thunderstorm+Hail"
	default:
		return "Unknown"
	}
}

// weatherCodeToEmoji returns a short symbol for weather
func weatherCodeToEmoji(code uint8) string {
	switch code {
	case 0:
		return "[Sun]"
	case 1, 2, 3:
		return "[Cloud]"
	case 45, 48:
		return "[Fog]"
	case 51, 53, 55, 61, 63, 65, 80, 81, 82:
		return "[Rain]"
	case 56, 57, 66, 67:
		return "[Ice]"
	case 71, 73, 75, 77, 85, 86:
		return "[Snow]"
	case 95, 96, 99:
		return "[Storm]"
	default:
		return ""
	}
}

// itoa converts int to string (simple implementation for embedded)
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var digits [12]byte
	i := len(digits)
	for n > 0 {
		i--
		digits[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		digits[i] = '-'
	}
	return string(digits[i:])
}

# ESP32-S3 Weather Bridge for Gopher Badge

This Arduino sketch runs on an Adafruit QT Py ESP32-S3 connected to the Gopher Badge via a Qwiic cable. It fetches weather data from the Open-Meteo API over WiFi and serves it to the Gopher Badge via I2C.

## Hardware Requirements

- Adafruit QT Py ESP32-S3
- Qwiic/STEMMA QT cable
- Gopher Badge

## Dependencies

Install the following libraries in Arduino IDE via **Sketch → Include Library → Manage Libraries...**:

1. **ArduinoJson** by Benoit Blanchon
2. **Adafruit NeoPixel** by Adafruit

The following libraries are included with the ESP32 board support:

- WiFi
- HTTPClient
- Wire

## Board Setup

1. Open **Tools → Board → Boards Manager...**
2. Search for **esp32** and install **esp32 by Espressif Systems**
3. Select **Tools → Board → esp32 → Adafruit QT Py ESP32-S3 No PSRAM**
4. Configure the following under **Tools**:
   - **USB CDC On Boot**: Enabled
   - **USB Mode**: Hardware CDC and JTAG
   - **Partition Scheme**: TinyUF2 4MB No OTA

## Configuration

Copy the credentials template and edit with your WiFi details:

```bash
cd esp32_weather_bridge
cp credentials.h.example credentials.h
# Edit credentials.h with your WiFi SSID and password
```

The `credentials.h` file is in `.gitignore` to prevent accidentally committing your WiFi credentials.

If you forget to create the file or leave the SSID empty, the NeoPixel will blink red and an error will be printed to the serial console.

## Building and Flashing

1. Connect the QT Py ESP32-S3 to your computer via USB-C
2. Select the correct port: **Tools → Port → /dev/cu.usbmodem...** (macOS) or **COM...** (Windows)
3. Click **Upload** (or press Cmd+U / Ctrl+U)
4. Open **Tools → Serial Monitor** (115200 baud) to verify it's working

You should see output like:

```
ESP32-S3 Weather Bridge starting...
I2C slave initialized at address 0x42 (SDA=41, SCL=40)
Connecting to WiFi: <YOUR_SSID>
....
WiFi connected!
IP address: 192.168.x.x
Fetching weather for 37.7749, -122.4194...
Weather data received
Current: 17°C, Code: 2
Today: 17/8°C
Putting WiFi to sleep...
```

## Usage

1. Disconnect the QT Py from USB
2. Connect it to the Gopher Badge via the Qwiic cable
3. Power on the Gopher Badge
4. Navigate to **Weather** in the menu
5. Select a location to view the weather

## NeoPixel Indicator

- **Green**: WiFi is active (connecting or fetching data)
- **Off**: WiFi is sleeping (between fetches, every 5 minutes)
- **Blinking red**: WiFi SSID not configured (edit the sketch)

## I2C Protocol

- **Address**: 0x42
- **Pins**: SDA=GPIO41, SCL=GPIO40 (STEMMA QT connector)

The Gopher Badge can send a location command (9 bytes) to change the weather location:

- Byte 0: Command (0x01 = set location)
- Bytes 1-4: Latitude (float32, little-endian)
- Bytes 5-8: Longitude (float32, little-endian)

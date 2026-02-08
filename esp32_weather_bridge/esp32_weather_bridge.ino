/*
 * ESP32-S3 Weather Bridge for Gopher Badge
 *
 * This sketch runs on an Adafruit QT Py ESP32-S3 connected via Qwiic cable.
 * It fetches weather data from Open-Meteo API and serves it via I2C slave.
 *
 * I2C Address: 0x42
 * Qwiic pins on QT Py ESP32-S3: SDA1 = GPIO41, SCL1 = GPIO40
 *
 * Default location: San Francisco (37.7749°N, 122.4194°W)
 * Supports dynamic location changes via I2C commands from GopherBadge.
 */

#include <WiFi.h>
#include <HTTPClient.h>
#include <Wire.h>
#include <ArduinoJson.h>
#include <Adafruit_NeoPixel.h>

// WiFi credentials - copy credentials.h.example to credentials.h and edit
#include "credentials.h"
const char *ssid = WIFI_SSID;
const char *password = WIFI_PASSWORD;

// I2C slave address
#define I2C_SLAVE_ADDR 0x42

// I2C command codes
#define CMD_SET_LOCATION 0x01

// QT Py ESP32-S3 STEMMA QT / Qwiic connector pins
// These are on the SECOND I2C bus (Wire1), NOT the default Wire pins!
// See: https://learn.adafruit.com/adafruit-qt-py-esp32-s3/pinouts
#define SDA_PIN 41 // SDA1 on STEMMA QT connector
#define SCL_PIN 40 // SCL1 on STEMMA QT connector

// NeoPixel configuration
#define NEOPIXEL_PIN 39
#define NEOPIXEL_POWER 38
Adafruit_NeoPixel pixel(1, NEOPIXEL_PIN, NEO_GRB + NEO_KHZ800);

// Current location (default to San Francisco: 37.7749°N, 122.4194°W)
float currentLat = 37.7749;
float currentLon = -122.4194;
volatile bool locationChanged = false;

// Weather data structure (packed for I2C transfer)
// Total: 32 bytes
struct WeatherData
{
  int8_t currentTemp;   // Current temperature in °C (1 byte)
  uint8_t currentCode;  // WMO weather code (1 byte)
  uint8_t windSpeed;    // Wind speed in km/h (1 byte)
  int8_t todayHigh;     // Today's high (1 byte)
  int8_t todayLow;      // Today's low (1 byte)
  uint8_t todayPrecip;  // Today's precipitation in mm (1 byte)
  uint8_t todayCode;    // Today's weather code (1 byte)
  int8_t day1High;      // Tomorrow's high (1 byte)
  int8_t day1Low;       // Tomorrow's low (1 byte)
  uint8_t day1Code;     // Tomorrow's weather code (1 byte)
  int8_t day2High;      // Day after tomorrow's high (1 byte)
  int8_t day2Low;       // Day after tomorrow's low (1 byte)
  uint8_t day2Code;     // Day after tomorrow's weather code (1 byte)
  uint8_t status;       // 0=no data, 1=valid, 2=error (1 byte)
  uint8_t lastUpdate;   // Minutes since last update (1 byte)
  uint8_t reserved[17]; // Reserved for future use (17 bytes)
} __attribute__((packed));

WeatherData weatherData;
volatile uint8_t i2cRegister = 0;
unsigned long lastFetchTime = 0;
const unsigned long fetchInterval = 5 * 60 * 1000; // 5 minutes

// Forward declarations
extern volatile uint32_t i2cRequestCount;
void wakeWiFi();

// Build weather URL dynamically based on current location
String buildWeatherUrl()
{
  String url = "https://api.open-meteo.com/v1/forecast?";
  url += "latitude=" + String(currentLat, 4);
  url += "&longitude=" + String(currentLon, 4);
  url += "&current=temperature_2m,weather_code,wind_speed_10m";
  url += "&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,weather_code";
  url += "&timezone=auto&forecast_days=3";
  return url;
}

// I2C request handler - called when master reads data
void onRequest()
{
  // Write all 32 bytes of weather data
  Wire.write((uint8_t *)&weatherData, sizeof(weatherData));
  i2cRequestCount++;
}

// I2C receive handler - called when master writes data
void onReceive(int numBytes)
{
  if (numBytes >= 1)
  {
    uint8_t cmd = Wire.read();
    numBytes--;

    if (cmd == CMD_SET_LOCATION && numBytes == 8)
    {
      // Read latitude as float32 (little-endian)
      uint8_t latBytes[4];
      for (int i = 0; i < 4; i++)
      {
        latBytes[i] = Wire.read();
      }
      float newLat;
      memcpy(&newLat, latBytes, 4);

      // Read longitude as float32 (little-endian)
      uint8_t lonBytes[4];
      for (int i = 0; i < 4; i++)
      {
        lonBytes[i] = Wire.read();
      }
      float newLon;
      memcpy(&newLon, lonBytes, 4);

      // Only update if location actually changed
      if (newLat != currentLat || newLon != currentLon)
      {
        currentLat = newLat;
        currentLon = newLon;
        locationChanged = true;
        Serial.printf("Location changed to: %.4f, %.4f\n", currentLat, currentLon);
      }
      return;
    }
  }

  // Discard any remaining bytes
  while (Wire.available())
  {
    Wire.read();
  }
}

void setup()
{
  Serial.begin(115200);
  delay(1000);
  Serial.println("ESP32-S3 Weather Bridge starting...");

  // Initialize NeoPixel (off initially)
  pinMode(NEOPIXEL_POWER, OUTPUT);
  digitalWrite(NEOPIXEL_POWER, HIGH); // Enable NeoPixel power
  pixel.begin();
  pixel.setPixelColor(0, 0, 0, 0); // Off
  pixel.show();

  // Initialize weather data
  memset(&weatherData, 0, sizeof(weatherData));
  weatherData.status = 0; // No data yet

  // Configure I2C pins with internal pull-ups before initializing
  pinMode(SDA_PIN, INPUT_PULLUP);
  pinMode(SCL_PIN, INPUT_PULLUP);
  delay(100);

  // Initialize I2C as slave
  // Note: On ESP32-S3, Wire.begin() with address enables slave mode
  // Using slower speed (50kHz) for better reliability
  bool success = Wire.begin(I2C_SLAVE_ADDR, SDA_PIN, SCL_PIN, 50000);
  if (success)
  {
    Wire.onRequest(onRequest);
    Wire.onReceive(onReceive);
    Serial.printf("I2C slave initialized at address 0x%02X (SDA=%d, SCL=%d)\n", I2C_SLAVE_ADDR, SDA_PIN, SCL_PIN);
  }
  else
  {
    Serial.println("ERROR: Failed to initialize I2C slave!");
  }

  // Check WiFi credentials
  if (strlen(ssid) == 0 || strcmp(ssid, "YOUR_WIFI_SSID") == 0)
  {
    Serial.println("ERROR: WiFi SSID not configured! Copy credentials.h.example to credentials.h and edit it.");
    // Blink NeoPixel red to indicate error
    while (true)
    {
      pixel.setPixelColor(0, pixel.Color(255, 0, 0));
      pixel.show();
      delay(500);
      pixel.setPixelColor(0, 0, 0, 0);
      pixel.show();
      delay(500);
    }
  }

  // Connect to WiFi
  Serial.printf("Connecting to WiFi: %s\n", ssid);

  // Turn NeoPixel green to indicate WiFi active
  pixel.setPixelColor(0, pixel.Color(0, 255, 0));
  pixel.show();

  WiFi.begin(ssid, password);

  int attempts = 0;
  while (WiFi.status() != WL_CONNECTED && attempts < 30)
  {
    delay(500);
    Serial.print(".");
    attempts++;
  }

  if (WiFi.status() == WL_CONNECTED)
  {
    Serial.println("\nWiFi connected!");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());

    // Fetch weather immediately for default location
    fetchWeather();

    // Put WiFi to sleep until next update
    sleepWiFi();
  }
  else
  {
    Serial.println("\nWiFi connection failed!");
    weatherData.status = 2; // Error

    // Turn NeoPixel off on failure
    pixel.setPixelColor(0, 0, 0, 0);
    pixel.show();
  }
}

// Counter for I2C activity debugging
volatile uint32_t i2cRequestCount = 0;

// WiFi sleep state
bool wifiAsleep = false;

// Wake WiFi from sleep
void wakeWiFi()
{
  if (wifiAsleep)
  {
    Serial.println("Waking WiFi...");

    // Turn NeoPixel green to indicate WiFi active
    pixel.setPixelColor(0, pixel.Color(0, 255, 0));
    pixel.show();

    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, password);

    int attempts = 0;
    while (WiFi.status() != WL_CONNECTED && attempts < 20)
    {
      delay(250);
      Serial.print(".");
      attempts++;
    }

    if (WiFi.status() == WL_CONNECTED)
    {
      Serial.println("\nWiFi reconnected!");
      wifiAsleep = false;
    }
    else
    {
      Serial.println("\nWiFi reconnect failed!");
    }
  }
}

// Put WiFi to sleep to save power
void sleepWiFi()
{
  if (!wifiAsleep && WiFi.status() == WL_CONNECTED)
  {
    Serial.println("Putting WiFi to sleep...");
    WiFi.disconnect(true);
    WiFi.mode(WIFI_OFF);
    wifiAsleep = true;

    // Turn NeoPixel off when WiFi is sleeping
    pixel.setPixelColor(0, 0, 0, 0);
    pixel.show();
  }
}

void loop()
{
  static unsigned long lastDebugTime = 0;

  unsigned long now = millis();
  bool needFetch = false;

  // Check if location changed (triggered by I2C command)
  if (locationChanged)
  {
    locationChanged = false;
    needFetch = true;
    Serial.printf("Fetching weather for new location: %.4f, %.4f\n",
                  currentLat, currentLon);
  }

  // Check if periodic update is needed (every 5 minutes)
  if (now - lastFetchTime >= fetchInterval)
  {
    needFetch = true;
  }

  // Fetch weather if needed
  if (needFetch)
  {
    // Wake WiFi if sleeping
    if (wifiAsleep)
    {
      wakeWiFi();
    }

    if (WiFi.status() == WL_CONNECTED)
    {
      fetchWeather();
      // Put WiFi to sleep after successful fetch
      sleepWiFi();
    }
    else if (!wifiAsleep)
    {
      // Try to reconnect
      Serial.println("WiFi disconnected, reconnecting...");
      WiFi.reconnect();
      delay(5000);
      if (WiFi.status() == WL_CONNECTED)
      {
        fetchWeather();
        sleepWiFi();
      }
    }
  }

  // Update minutes since last update
  if (weatherData.status == 1)
  {
    weatherData.lastUpdate = min(255UL, (now - lastFetchTime) / 60000);
  }

  // Debug output every 10 seconds
  if (now - lastDebugTime >= 10000)
  {
    Serial.printf("[DEBUG] I2C requests: %lu, Status: %d, Temp: %d°C, Location: %.4f,%.4f, WiFi: %s\n",
                  i2cRequestCount, weatherData.status, weatherData.currentTemp,
                  currentLat, currentLon,
                  wifiAsleep ? "sleeping" : "active");
    lastDebugTime = now;
  }

  delay(100);
}

void fetchWeather()
{
  Serial.printf("Fetching weather for %.4f, %.4f...\n",
                currentLat, currentLon);

  String url = buildWeatherUrl();
  Serial.printf("URL: %s\n", url.c_str());

  HTTPClient http;
  http.begin(url);
  http.setTimeout(10000);

  int httpCode = http.GET();

  if (httpCode == HTTP_CODE_OK)
  {
    String payload = http.getString();
    Serial.println("Weather data received");

    // Parse JSON
    JsonDocument doc;
    DeserializationError error = deserializeJson(doc, payload);

    if (!error)
    {
      // Current weather
      weatherData.currentTemp = (int8_t)doc["current"]["temperature_2m"].as<float>();
      weatherData.currentCode = doc["current"]["weather_code"].as<uint8_t>();
      weatherData.windSpeed = (uint8_t)doc["current"]["wind_speed_10m"].as<float>();

      // Daily forecast (today, tomorrow, day after)
      JsonArray maxTemps = doc["daily"]["temperature_2m_max"];
      JsonArray minTemps = doc["daily"]["temperature_2m_min"];
      JsonArray precip = doc["daily"]["precipitation_sum"];
      JsonArray codes = doc["daily"]["weather_code"];

      if (maxTemps.size() >= 3)
      {
        weatherData.todayHigh = (int8_t)maxTemps[0].as<float>();
        weatherData.todayLow = (int8_t)minTemps[0].as<float>();
        weatherData.todayPrecip = (uint8_t)min(255.0f, precip[0].as<float>());
        weatherData.todayCode = codes[0].as<uint8_t>();

        weatherData.day1High = (int8_t)maxTemps[1].as<float>();
        weatherData.day1Low = (int8_t)minTemps[1].as<float>();
        weatherData.day1Code = codes[1].as<uint8_t>();

        weatherData.day2High = (int8_t)maxTemps[2].as<float>();
        weatherData.day2Low = (int8_t)minTemps[2].as<float>();
        weatherData.day2Code = codes[2].as<uint8_t>();
      }

      weatherData.status = 1; // Valid data
      lastFetchTime = millis();

      Serial.printf("Current: %d°C, Code: %d\n", weatherData.currentTemp, weatherData.currentCode);
      Serial.printf("Today: %d/%d°C\n", weatherData.todayHigh, weatherData.todayLow);
    }
    else
    {
      Serial.printf("JSON parse error: %s\n", error.c_str());
      weatherData.status = 2; // Error
    }
  }
  else
  {
    Serial.printf("HTTP error: %d\n", httpCode);
    weatherData.status = 2; // Error
  }

  http.end();
}

# Gopher Badge v1.2 - Full Technical Databrief

## Page 1: Features and Introduction

**Tiny GO Badge Board Data brief** **Purple Hardware by Pólux**

### Features

**Common features**

- RP2040 MCU, dual core M0+
- 64Mbit of Flash Quad-SPI
- 1 user Led
- Accel LIS3DHTR sensor
- 2 RGB serial Neopixel compatible
- 6 user buttons
- 12Mhz Oscillator
- IPS TFT Display of 320x240p
- Buzzer and amplifier for sound reproduction

**Board connections:**

- SWDIO
- USB C, battery charge and data
- Qwiic/Stemma
- Solder Pads accessible for expansion purposes

**Power:**

- Work with LiPo and/or with the USB C
- USB C LiPo battery charging

**Community Support of several languages:**

- Tiny Go, C/C++, Arduino, Mbed, Micro Python

**Introductory Text:**
The tiny go badge is a board inspired by the Gopher pet of Goland in association with the GO community. This board is created with educational purposes focused on the future presentations around the world make by the GO community.

---

## Page 2: Pin Demux (Table 1)

**To understand how to work with this board the following table will introduce the GPIO and the functionality**. This support could be added by the community on the language supported for the board configuration or should be added manually by the user.

| Description                                                     | GPIO | Function    | Bus/PWM Channel | Voltage level |
| --------------------------------------------------------------- | ---- | ----------- | --------------- | ------------- |
| I2C on RP2040                                                   | 0    | SDA         | 0               | 3,3V          |
| I2C on RP2040                                                   | 1    | SCL         | 0               | 3,3V          |
| User Led                                                        | 2    | Output      | —               | 3,3V          |
| Amplifier enable control (High enable)                          | 3    | Output      | —               | 3,3V          |
| External acces by "Solder Pads"                                 | 4    | Not defined | —               | 3,3V          |
| External acces by "Solder Pads"                                 | 5    | Not defined | —               | 3,3V          |
| External acces by "Solder Pads"                                 | 6    | Not defined | —               | 3,3V          |
| External acces by "Solder Pads"                                 | 7    | Not defined | —               | 3,3V          |
| External acces by "Solder Pads"                                 | 8    | Not defined | —               | 3,3V          |
| Tearing effect display control                                  | 9    | Output      | —               | 3,3V          |
| Input A front button                                            | 10   | Input       | —               | 3,3V          |
| Input B front button                                            | 11   | Input       | —               | 3,3V          |
| Display Backlight enable (High enable, can be PWM or 0/1 logic) | 12   | Output      | Channel 6A      | 3,3V          |
| Not in use (PenIRQ with NS2009 resistive controller)            | 13   | Input       | —               | 3,3V          |
| Analog PMW Sound Output                                         | 14   | Output      | Channel 7A      | 3,3V          |
| Serial Bus for Neopixel type RGB Led                            | 15   | Serial      | —               | 3,3V          |
| TFT MISO                                                        | 16   | Output      | 0               | 3,3V          |
| TFT CS pin                                                      | 17   | Output      | —               | 3,3V          |
| TFT SPI CLK                                                     | 18   | Output      | 0               | 3,3V          |
| TFT SPI MOSI                                                    | 19   | Output      | 0               | 3,3V          |
| TFT DC pin                                                      | 20   | Output      | —               | 3,3V          |
| Reset display (Low reset)                                       | 21   | Output      | —               | 3,3V          |
| Input right front button                                        | 22   | Input       | —               | 3,3V          |
| Input down front button                                         | 23   | Input       | —               | 3,3V          |
| Input up front button                                           | 24   | Input       | —               | 3,3V          |
| Input left front button                                         | 25   | Input       | —               | 3,3V          |
| External acces by "Solder Pads"                                 | ADC0 | Not defined | —               | 3,3V          |
| External acces by "Solder Pads"                                 | ADC1 | Not defined | —               | 3,3V          |
| External acces by "Solder Pads"                                 | ADC2 | Not defined | —               | 3,3V          |
| Not in use                                                      | ADC3 | —           | —               | 3,3V          |

---

## Page 3: Power Section

**The board can work at the same time with battery and with the USB C connected**. If the USB C is connected and the board have a battery connected, this will start the charging process with a maximum charge rate of 500mA.

**Note:** ADC3 is not connected and there is no battery voltage divider, so it's not possible to read the battery voltage and estimate the battery level.

### Schematic Notes (Power)

- **DESIGN NOTE:** Slide Switch. Left VBAT(ON), Right.
- **DESIGN NOTE:** USB 500mA Charging.
- **Components:**
- **SW9:** SPDT Switch.
- **D1:** 0.31VF 1A Diode.
- **U6 (Regulator):** VIN VOUT, EN GND, 3.3V 0.6A 320mV.
- **U4 (Charger):** VDD, STAT, VSS, VBAT, PROG (4.2V LiPo Charger).
- **R19:** 2K 1% 0603 (connected to PROG).
- **LED1:** RED LED 0603 (Charging status).
- **J4:** JST PH Connector.

---

## Page 4: USB Type C

**DESIGN NOTE:** CC1 & CC2 PullDown, always 5V compilant.

### Schematic Notes (USB)

- **LAYOUT NOTE:** USB Zdiff 90R Zr 50R.
- **DIFFPAIR:** USB LANES TYPEC (USB_TYPEC_P, USB_TYPEC_N).
- **Resistors:** R20, R21 (5.1K 1% 0603).
- **TVS1:** Protection diode.
- **FB1:** 120OHM 0603.

---

## Page 5: RP2040 MCU

The MCU of the board it's a RP2040 from raspberry pi. This MCU have a dual core M0+ capable to run up to 130Mhz. The Tiny Go badge have a 64mbit (8MegaByte) flash memory to store all program contents.

### Schematic Notes (MCU)

- **DESIGN NOTE:** Power consumption average 30mA.
- **LAYOUT NOTE:** USB Zdiff 90R Zr 50R.
- **Oscillator:** 12Mhz 20ppm.
- **Internal Pins:** SWDIO, SWCLK, USB DP, USB DM, RUN, XIN, XOUT.
- **Reference Voltage:** ADC AVDD, VREG VIN, VREG VOUT.

---

## Page 6: SPI Flash and Buttons

The SPI Flash content its accessible by the RP2040 bootloader if the Boot switch (SW7) is pressed at the reset of the board, this reset could be done by a momentary press of the Reset button (SW8).

### Schematic Notes (Flash/Reset)

- **U2:** W25Q16JVSSIQ (16Mbit 133Mhz 8SOIC).
- **SW7 (Boot):** Tactile Switch SKRPACE010.
- **SW8 (Reset):** Tactile Switch SKRPACE010.
- **R8:** 1K 1% 1/10W 0603.

---

## Page 7-8: Accelerometer (LIS3DHTR)

The tiny go badge board have an accelerometer sensor that is accessible by the I2C bus, this sensor it's a well know one by the community the LIS3DHTR. The address of the sensor is editable with solder tools, by default the R17 is soldered on the board, but if the user wants to change the I2C address, it could be done by desolder the R17 and solder the R16.

### Schematic Notes (Accelerometer)

- **DESIGN NOTE:** CS Pull-Up = I2C Interface.
- **DESIGN NOTE:** SDO Pull-Down = I2C address LSB to 0.
- **U3:** LIS3DHTR (Accel 2-16G).
- **Note:** The modification mentioned above its marked on the PCB board silkscreen.

---

## Page 9-10: User Inputs and Audio

**User buttons:** On the board are a total of 6 user buttons, they are allocated in the front. Those buttons have a logic Low activation.
**User Led and RGB:** A total of 2 RGB serial Led are allocated in the eyes of the Gopher, and one standard Led its allocated in the back of the board. The RGB LEDs are compatible with the NeoPixel library maintained by the community.
**Audio:** The tiny go Badge has the ability to play sounds with the help of a buzzer. This sounds must be created by the RP2040 and modulated by PWM.

### Schematic Notes (Input/Audio)

- **Buttons:** SW1, SW2, SW3, SW4, SW5, SW6 (Tactile Switch SKRPACE010).
- **RGB LEDs:** RGB1, RGB2 (LED 2020).
- **User LED:** LED2 (RED LED 0603).
- **Audio Amp (U5):** 621-PAM8302AADCR (2.5W 4R).
- **DESIGN NOTE:** Sound PWM Channel 7A GPIO RPI GPIO14.

---

## Page 11: External Connectors

This board its compatible with the Qwiic/Stema 3,3V logic connectors. In addition, the user have a Solder Pads exposed to expansion purposes on the board.

### Schematic Notes (Expansion)

- **DESIGN NOTE:** We should not exceed 200mA on 3.3V.
- **J2, J3, J7:** QWIIC/JST SH connectors.
- **Solder PAD:** 12-pin breakout.

---

## Page 12-14: Display Details

The tiny go Badge board have the ability to drive a SPI TFT display of 50 pins with backlight control. The backlight control could be implemented by a PWM or a logic 1/0. One extra in the board it's the capability of solder a NS2009 resistive touch controller, this controller isn't on the scope of this board, but if the user wants to buy an extra SPI TFT display with resistive touch, the board have the footprint to make it work.

### Schematic Notes (Display)

- **DESIGN NOTE (Interface):** IM3-1, IM2=1, IM1=1, IM0-0 4wire 8-bit data Serial Interface => SCL, SDI, DCX, SDO, CSX.
- **DESIGN NOTE (Backlight):** Backlight PWM Channel 6A GPIO 12. On default ON.
- **DESIGN NOTE (TE):** RPI GPIO9 TE without use.
- **Control Pins:** RST GPIO 21, GP20= DC, GP17 = CS, GP18 SPICLK, GP19 SPIOTX, GP16SPIORX.
- **U7:** Resistive Touch Controller NS2009 (Footprint on design, but not mounted).

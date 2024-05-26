# Titan W Firmware

This firmware will enable Titan W to expose accelerometer data using Bluetooth LE, notification and GATT read mode are both supported.

Bluetooth Service UUID: `E0E62F49-9E34-462C-8240-720C01FC5374`

Acceleration Characteristic UUID: `DC77CEBA-DEE1-4467-9814-E2EA1DA1098A`

## Hardware requirements

### SoC requirement

Current code is targeting nRF52840 SoC, and should be able to run on any [Zephyr Supported Boards](https://docs.zephyrproject.org/latest/boards/index.html#boards) with Bluetooth support.

### Accelerometer requirements

We currently support using [LIS3DH](https://www.adafruit.com/product/2809) as the accelerometer, over I2C bus. The SDA pin should be connected to P0.29, and SCL should be connected to P0.31.

## Build Instructions

1. Use nRF Connect to install the latest toolchain, and VS Code plugins.
2. Open this repo on VS Code, go to the nRF Connect side panel, and create a build configuration:
    - Select `nrf52840dongle_nrf52840` as the board
    - Add [boards/titan_w.overlay](boards/titan_w.overlay) as the `Devicetree Overlays`
3. Build using the configuration
4. The artifact should be available at `build/zephyr/zephyr.hex`

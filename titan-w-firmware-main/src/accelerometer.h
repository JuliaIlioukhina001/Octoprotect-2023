#ifndef ACCELEROMETER_H
#define ACCELEROMETER_H

#include <zephyr/types.h>
#include <zephyr/drivers/i2c.h>

// Initialize the accelerometer
void accelerometer_init();

// Read current acceleration data (6 bytes, XL|XH|YL|YH|ZL|ZH) into the given buffer
void accelerometer_read(uint8_t *data);

#define LIS3DH_WHOAMI_CONSTANT   0b00110011

#define LIS3DH_REBOOT            0b10000000
#define LIS3DH_ENABLE_XYZ_AXIS   0b00000111
#define LIS3DH_DATA_RATE_10HZ    0b00100000
#define LIS3DH_DATA_RATE_100HZ   0b01010000
#define LIS3DH_BDU               0b10000000
#define LIS3DH_HI_RES            0b00001000
#define LIS3DH_ADDR_INCREMENT    0b10000000

#define LIS3DH_REG_WHO_AM_I      0x0f
#define LIS3DH_REG_TEMP_CFG      0x1f
#define LIS3DH_REG_CTRL1         0x20
#define LIS3DH_REG_CTRL2         0x21
#define LIS3DH_REG_CTRL3         0x22
#define LIS3DH_REG_CTRL4         0x23
#define LIS3DH_REG_CTRL5         0x24
#define LIS3DH_REG_CTRL6         0x25
#define LIS3DH_REG_OUT_X_L       0x28
#define LIS3DH_REG_OUT_X_H       0x29
#define LIS3DH_REG_OUT_Y_L       0x2a
#define LIS3DH_REG_OUT_Y_H       0x2b
#define LIS3DH_REG_OUT_Z_L       0x2c
#define LIS3DH_REG_OUT_Z_H       0x2d

#endif
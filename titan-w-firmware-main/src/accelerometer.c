#include "accelerometer.h"

#define I2C0_NODE DT_NODELABEL(accelerometer)

static const struct i2c_dt_spec dev_i2c = I2C_DT_SPEC_GET(I2C0_NODE);

void accelerometer_init(){
    if (!device_is_ready(dev_i2c.bus)) {
        printk("[ERROR] I2C bus %s is not ready!\n",dev_i2c.bus->name);
        return;
    }
    uint8_t buf;
    // Read WHOAMI register to make sure it's the right device
    i2c_burst_read_dt(&dev_i2c, LIS3DH_REG_WHO_AM_I, &buf, 1);
    if (buf != LIS3DH_WHOAMI_CONSTANT){
        printk("[ERROR] Device at 0x18 is not LIS3DH accelerometer. Value at LIS3DH_REG_WHO_AM_I(0x0F) is %x!\n", buf);
        return;
    }
    // Set data rate to 10Hz, and enable xyz-axis acceleration measurement
    buf = LIS3DH_DATA_RATE_10HZ | LIS3DH_ENABLE_XYZ_AXIS;
    i2c_burst_write_dt(&dev_i2c, LIS3DH_REG_CTRL1, &buf, 1);
    // Set the Block Data Update flag, and High Resolution flag
    buf = LIS3DH_BDU | LIS3DH_HI_RES;
    i2c_burst_write_dt(&dev_i2c, LIS3DH_REG_CTRL4, &buf, 1);
}

void accelerometer_read(uint8_t *data){
    // Read 6 bytes starting from OUT_X_L register
    int ret = i2c_burst_read_dt(&dev_i2c, LIS3DH_REG_OUT_X_L | LIS3DH_ADDR_INCREMENT, data, 6);
    if(ret != 0){
        printk("[ERROR] Failed to read from I2C device address %x!\n", dev_i2c.addr);
    }
}

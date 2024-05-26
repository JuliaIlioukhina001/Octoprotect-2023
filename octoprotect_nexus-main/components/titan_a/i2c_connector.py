import board
import adafruit_tca9548a


def get_all_connected() -> (None, int):
    i2c = board.I2C()
    # Create the TCA9548A object and give it the I2C bus
    tca = adafruit_tca9548a.TCA9548A(i2c)
    print("Searching for I2C Devices connected to the multiplexer...")
    for channel in range(8):
        if tca[channel].try_lock():
            print("Channel {}:".format(channel), end="")
            addresses = tca[channel].scan()
            print([hex(address) for address in addresses if address != 0x70])
            tca[channel].unlock()
            # Check if the device is an accelerometer
            if 0x18 in addresses:
                yield tca[channel], channel


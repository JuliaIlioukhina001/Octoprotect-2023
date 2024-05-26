import math
import random
import uuid

import adafruit_lis3dh

from components.model.accel_data_source import AccelerometerDataSource


class I2CDataSource(AccelerometerDataSource):
    def __init__(self, index: int, channel):
        self.pz = 0
        self.py = 0
        self.px = 0
        rd = random.Random()
        rd.seed(index)
        self.accel = adafruit_lis3dh.LIS3DH_I2C(channel)
        self.accel.data_rate = adafruit_lis3dh.DATARATE_10_HZ
        self.unique_hash = uuid.UUID(int=rd.getrandbits(128))
        self.working = True

    def get_accelerometer_name(self) -> str:
        """Returns the friendly name of the type of accelerometer"""
        return "Titan A - Attached Accelerometer"

    def get_magnitude_change(self) -> float:
        """Returns the change in magnitude of acceleration in m/s^2 since the last update"""
        try:
            """Returns the magnitude of acceleration in m/s^2"""
            x, y, z = self.accel.acceleration
            mag = math.sqrt((x - self.px) ** 2 + (y - self.py) ** 2 + (z - self.pz) ** 2)
            self.px = x
            self.py = y
            self.pz = z
            return mag
        except:
            self.working = False

    def get_unique_id(self) -> uuid:
        """Gets a universal id for the sensor"""
        return self.unique_hash

    def stop(self):
        pass

    def is_working(self) -> bool:
        return self.working

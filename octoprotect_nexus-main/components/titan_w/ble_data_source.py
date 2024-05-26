import asyncio.tasks
import hashlib
import math
import uuid
from asyncio import Future

from bleak import BleakClient, BLEDevice

import struct
from typing import Tuple
from components.model.accel_data_source import AccelerometerDataSource
from components.titan_w.bt_tw_constants import ACCEL_CHAR, TITAN_BLE_SERVICE

RANGE_2_G = 0b00
STANDARD_GRAVITY = 9.806


def acceleration(data) -> Tuple[float, float, float]:
    """The x, y, z acceleration values returned
    in a 3-tuple and are in :math:`m / s ^ 2`"""
    divider = 16380

    x, y, z = struct.unpack("<hhh", data)

    # convert from Gs to m / s ^ 2 and adjust for the range
    x = (x / divider) * STANDARD_GRAVITY
    y = (y / divider) * STANDARD_GRAVITY
    z = (z / divider) * STANDARD_GRAVITY
    return x, y, z


class BLEDataSource(AccelerometerDataSource):
    future: Future

    async def _device_connection(self, device):
        """
        Listens to the acceleration updates from this device
        """
        try:
            async with BleakClient(
                    device,
                    services=[TITAN_BLE_SERVICE]
            ) as client:
                await client.pair()

                def callback(_, arr: bytearray):
                    x, y, z = acceleration(arr)
                    self.mag = math.sqrt((x - self.px) ** 2 + (y - self.py) ** 2 + (z - self.pz) ** 2)
                    self.px = x
                    self.py = y
                    self.pz = z

                await client.start_notify(ACCEL_CHAR, callback)
                while not self.future.done():
                    await asyncio.sleep(0.1)
                    # if not client.is_connected:
                    #     break
                await client.disconnect()
        except Exception as e:
            print(f"Unhandled exception in BLE Data Source: {e}")
        self.connected = False

    def __init__(self, device: BLEDevice):
        self.px = 0
        self.py = 0
        self.pz = 0
        self.mag = 0
        self.future = Future()
        self.connected = True
        hex_string = hashlib.md5(device.address.encode("UTF-8")).hexdigest()
        self.unique_hash = uuid.UUID(hex=hex_string)
        asyncio.tasks.ensure_future(self._device_connection(device))

    def get_accelerometer_name(self) -> str:
        """Returns the friendly name of the type of accelerometer"""
        return "Titan W - Wireless Accelerometer"

    def get_magnitude_change(self) -> float:
        """Returns the change in magnitude of acceleration in m/s^2 since the last update"""
        return self.mag

    def get_unique_id(self) -> uuid:
        """Gets a universal id for the sensor"""
        return self.unique_hash

    def stop(self):
        self.future.set_result(None)

    def is_working(self) -> bool:
        return self.connected

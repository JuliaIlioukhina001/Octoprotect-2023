import hashlib

from bleak import BleakScanner, BleakClient, AdvertisementData, BLEDevice

from components.titan_w.bt_tw_constants import *


async def scan():
    """
    Searches all BLE advertisements for Titan W advertisements. See Constants for the uuid.
    """
    print("Scanning for Titan W over BLE")
    devices = await BleakScanner.discover(return_adv=True)
    found_devices = []
    adv: AdvertisementData
    device: BLEDevice
    for mac in devices:
        device, adv = devices[mac]
        if TITAN_BLE_SERVICE in adv.service_uuids:
            hwid = hashlib.md5(device.address.encode("UTF-8")).hexdigest()
            print(f"Discovered device {adv.local_name} with hwid: {hwid}")
            async with BleakClient(
                    device,
                    services=[TITAN_BLE_SERVICE]
            ) as client:
                svc = client.services[TITAN_BLE_SERVICE]
                await client.pair()
                if ACCEL_CHAR in map(lambda x: x.uuid, svc.characteristics):
                    found_devices.append(device)
    return found_devices

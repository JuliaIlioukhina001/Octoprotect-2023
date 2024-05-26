import asyncio
import hashlib
import json
import uuid
from asyncio import Future

from components.backend.server_link import ServerLink
from components.model.accel_data_source import AccelerometerDataSource
from components.model.state_service import NexusState
from components.services.accel_monitor import AcceleratorMonitor
from components.titan_w.ble_data_source import BLEDataSource
from components.titan_w.ble_searcher import scan

Monitor: AcceleratorMonitor
Server: ServerLink
State: NexusState


async def initialize(error):
    """
    Initializes and creates all the singleton instances of services
    """
    global Monitor, Server, State

    print("Initializing Components...")

    Monitor = AcceleratorMonitor(5)

    print("Connecting to the backend server...")
    config = json.load(open("config.json", 'r'))

    Server = ServerLink(config['url'])

    State = NexusState()
    State.error = error

    await Server.connect(config['token'])


async def init_devices():
    """
    Scans (and reconnects to pre-existing) devices over I2C and BLE.
    """
    from components.alarm.alarm_service import clear_alarm
    conn: [AccelerometerDataSource]
    State.completion.set_result(None)  # stop existing threads
    clear_alarm()
    for conn in Monitor.sources:
        conn.stop()
        await Server.send_connection_state(conn.get_unique_id(), False)

    Monitor.sources = []
    State.completion = Future()

    try:
        # lazy import so we don't get errors
        from components.titan_a.i2c_connector import get_all_connected
        from components.titan_a.i2c_data_source import I2CDataSource
        # Add all TITAN A accelerometers
        for i2c, channel in get_all_connected():
            source = I2CDataSource(channel, i2c)
            Monitor.add_data_source(source)
            await Server.send_connection_state(source.get_unique_id(), True)
    except NotImplementedError as ex:
        print(f"Error while initializing TITAN A Monitors {ex}")

    for device in await scan():
        hwid = hashlib.md5(device.address.encode("UTF-8")).hexdigest()
        if str(hwid) in State.paired:
            Monitor.add_data_source(BLEDataSource(device))
            await Server.send_connection_state(uuid.UUID(hex=hwid), True)

    asyncio.ensure_future(Monitor.monitor(State.completion))

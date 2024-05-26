import asyncio
import json

import websockets
from websockets import WebSocketClientProtocol


class ServerLink:
    """
    Establishes a real-time connection with the backend server
    """
    socket: WebSocketClientProtocol

    def __init__(self, socket_addr):
        self.host = socket_addr

    async def connect(self, token):
        """
        Connects to the backend server, and subscribes to events
        """
        self.socket = await websockets.connect(self.host,
                                               extra_headers={"Authorization": f"Bearer {token}"})
        asyncio.ensure_future(self._listen())

    async def _listen(self):
        from components.services.service_registry import State, init_devices
        from components.alarm.alarm_service import clear_alarm
        # Handles events sent by the server
        print("Listening to server events:")
        try:
            while True:
                msg = await self.socket.recv()
                data = json.loads(msg)
                t = data['type']
                print(f"Got packet {t}")
                if t == 'arm':
                    State.is_armed = True
                    print("Device Armed.")
                elif t == 'disarm':
                    State.is_armed = False
                    clear_alarm()
                    print("Device Disarmed.")
                elif t == 'start-stream':
                    State.stream = True
                    print("Streaming acceleration data.")
                elif t == 'stop-stream':
                    State.stream = False
                    print("Streaming stopped.")
                elif t == 'initialize':
                    State.paired = data["devices"]
                    State.sensitivity = data["sensitivity"]
                    print(f"Pairing with devices: {State.paired}")
                    await init_devices()
                elif t == 'request-state':
                    await self.send_nexus_state()
                    print("Sent state to backend.")
        except Exception as e:
            print(f'Websocket Exception {e}')
            State.error.set_result(None)

    async def send_accel(self, uuid, mag):
        """
        Sends the acceleration data in the streaming state
        :param uuid: The UUID of the Titan
        :param mag: The magnitude in m/s^2
        """
        await self.socket.send(json.dumps({
            'type': 'accel',
            'titanID': str(uuid),
            'magnitude': mag
        }))

    async def send_trigger(self, uuid, mag):
        """
        Sends the acceleration data, and signals the backend that an alarm is triggered
        :param uuid: The UUID of the Titan
        :param mag: The magnitude in m/s^2
        """
        await self.socket.send(json.dumps({
            'type': 'movement-trigger',
            'titanID': str(uuid),
            'magnitude': mag
        }))

    async def send_connection_state(self, uuid, is_connected):
        """
        Sends a packet to the server notifying of a device connecting/disconnecting from the Nexus
        :param uuid: The UUID of the Titan
        :param is_connected: The new connection state of the device
        :return:
        """
        await self.socket.send(json.dumps({
            'type': 'conn-state',
            'titanID': str(uuid),
            'isConnected': is_connected
        }))

    async def send_nexus_state(self):
        """
        Reports the current state of the nexus to the backend
        """
        from components.services.service_registry import State, Monitor
        await self.socket.send(json.dumps({
            'type': 'nexus-state',
            'titan': [
                {
                    'id': str(x.get_unique_id()),
                    'name': x.get_accelerometer_name(),
                    'isWorking': x.is_working()
                } for x in Monitor.sources
            ],
            'isArmed': State.is_armed,
            'isTriggered': State.triggered
        }))

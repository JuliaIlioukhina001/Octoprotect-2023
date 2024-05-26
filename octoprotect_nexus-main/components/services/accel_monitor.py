import asyncio
import time
from asyncio import Future

from components.alarm.alarm import stop_alarm, play_alarm
from components.alarm.alarm_service import trigger_alarm
from components.model.accel_data_source import AccelerometerDataSource


class AcceleratorMonitor:
    sources: [AccelerometerDataSource]

    def __init__(self, arm_delay):
        """
        Initializes the acceleration monitor
        :param arm_delay: the number of seconds after which the monitor outputs data
        """
        self.delay = arm_delay
        self.sources = []

    def add_data_source(self, source: AccelerometerDataSource):
        """
        Adds a datasource for the monitor to pull
        :param source: the implementation
        """
        self.sources.append(source)

    async def monitor(self, cancellation: Future):
        """
        Begins monitoring for movements. The Nexus MUST be already connected to the backend, since it forwards the
        acceleration via ws
        """
        from components.services.service_registry import Server, State
        print("Monitoring...")
        start = time.time()
        while not cancellation.done():
            source: AccelerometerDataSource
            for source in self.sources:
                if not source.is_working():
                    continue

                mag = source.get_magnitude_change()
                if State.stream:
                    await Server.send_accel(source.get_unique_id(), mag)

                if State.is_armed and mag >= State.sensitivity and time.time() - start > self.delay:
                    print(f"Movement Detected, {source.get_accelerometer_name()} id - {source.get_unique_id()} {mag}")
                    trigger_alarm()
                    await Server.send_trigger(source.get_unique_id(), mag)

                if not source.is_working():
                    print(f'Failed while reading data source {source.get_accelerometer_name()}. Warning the backend...')
                    await Server.send_connection_state(source.get_unique_id(), False)

            await asyncio.sleep(0.1)

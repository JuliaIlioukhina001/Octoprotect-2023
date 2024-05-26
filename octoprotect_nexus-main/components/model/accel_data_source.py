import uuid


class AccelerometerDataSource:
    def get_accelerometer_name(self) -> str:
        """Returns the friendly name of the type of accelerometer"""
        pass

    def get_magnitude_change(self) -> float:
        """Returns the change in magnitude of acceleration in m/s^2 since the last update"""
        pass

    def get_unique_id(self) -> uuid:
        """Gets a universal id for the sensor"""
        pass

    def stop(self):
        """Disconnects from the data source"""
        pass

    def is_working(self) -> bool:
        pass

# OctoProtect

OctoProtect is an anti-theft system in the form of an octopus, created by Adam Chen, Paul Lee, Dongruixuan Li, Julia Ilioukhina, Mariya Turetska, and Jane Zeng.

## Summary

OctoProtect protects personal belongings using the NEXUS device with tentacles and wireless TITANs. TITANs contain accelerometers to detect movement, triggering notifications and alarms with LEDs and buzzers. Ideal for university students in public spaces.

## Technology Stack

- **Frontend:** React Native, Redux, Material UI
- **Backend:** Golang, Gin, GORM, WebSocket, JWT, Bcrypt
- **Firmware:** Nordic nRF52840, Zephyr RTOS, C
- **Deployment:** Docker, GitLab CI/CD, Kubernetes

## Features

- **Detection:** Accelerometers detect movement.
- **Notifications:** Alerts via mobile app and built-in alarm.
- **Interactive App:** Built with React Native, Redux, and WebSocket.
- **Secure Backend:** Uses JWT and Bcrypt for authentication and security.
- **Firmware:** Low latency communication with Nordic nRF52840 and Zephyr RTOS.

## Key Components

- **NEXUS:** Central hub running Python service, connects to Backend and TITANs.
- **TITAN W:** Wireless devices using BLE for communication.
- **Hardware:** Raspberry Pi Zero W 2, I2C Multiplexer, Adafruit LIS3DH Accelerometer.

## Intellectual Property

Used open-source libraries under MIT, BSD, Apache, and MPL licenses.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

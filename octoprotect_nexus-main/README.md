# Octoprotect Nexus

The Nexus software orchestrates the interop between the backend server and the Titan W and Titan A hardware components.

#### Its key goals are to:

- Facilitate a realtime Websocket connection with the Backend
- Monitor Acceleration over:
  - I2C Multiplexer - (`Titan A`)
  - Bluetooth Low Energy - (`Titan W`)
- Trigger an audio alarm when the magnitude of acceleration reaches a threshold.

# Installation

1. Clone the repository into your home folder and name it `nexus`.

2. Modify `config.json` with the relevant information.

3. Install Nexus as a `systemd` service with the following configuration:

```ini
[Unit]
Description=Nexus Host Service
After=multi-user.target
[Service]
Type=simple
Restart=always
ExecStart=/usr/bin/python3 /home/<username>/nexus/main.py
WorkingDirectory=/home/<username>/nexus
[Install]
WantedBy=multi-user.target
```

and save it to `/etc/systemd/system/nexus.service`

4. Run the following to complete the setup:

```shell
sudo systemctl daemon-reload
sudo systemctl enable nexus.service
sudo systemctl start nexus.service
```
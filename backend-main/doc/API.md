# Octoprotect Backend API Documentation

## Websocket Endpoint

### Convention

Each message has the following structure:

```json
{
  "type": "<message type>",
  ...
}
```

Error message has the following structure:

```json
{
  "type": "error",
  "message": "<error message>"
}
```

For detailed types of messages below, [C] means the message should be sent by client, and [S] means the message should be sent by server. 

### Nexus Side

Endpoint: `wss://octoprotect.starcatmeow.cn/ws/nexus`

Authentication: Put the signed JWT with `nexusMac` claim in Authorization header, with prefix `Bearer `

Messages may contain `nexusID` field, can safely ignore it.

#### [S] Type: `initialize`

Server will send this message when Nexus connects to server, and when user updates new config. 

```json
{
    "type": "initialize",
    "devices": [ // list of enabled Titan W
      "EC35CF7C-E24F-4311-8BA3-5002B90C04D7",
      ...
    ],
    "sensitivity": 1.0,
}
```

#### [S] Type: `request-state`

Server will send this message when user requests state. Nexus should respond with `nexus-state`.

```json
{
  "type": "request-state"
}
```

#### [C] Type: `nexus-state`

Nexus should send this message as response of `request-state`.

```json
{
  "type": "nexus-state",
  "titan": [
    {
      "id": "e3e70682-c209-4cac-629f-6fbed82c07cd",
      "name": "Titan A - Attached Accelerometer",
      "isWorking": true
    },
    {
      "id": "cd613e30-d8f1-6adf-91b7-584a2265b1f5",
      "name": "Titan A - Attached Accelerometer",
      "isWorking": true
    },
    {
      "id": "6513270e-269e-0d37-f2a7-4de452e6b438",
      "name": "Titan A - Attached Accelerometer",
      "isWorking": true
    }
  ],
  "isArmed": false
}
```

#### [C] Type: `conn-state`

Nexus should send this message when there is a Titan connection change.

```json
{
  "type": "conn-state",
  "titanID": "6513270e-269e-0d37-f2a7-4de452e6b438",
  "isConnected": false
}
```

#### [S] Type: `start-stream`

Server will send this message to enable Nexus acceleration data streaming. Nexus should then start sending `accel` messages.

```json
{
  "type": "start-stream"
}
```

#### [S] Type: `stop-stream`

Server will send this message to disable Nexus acceleration data streaming. Nexus should then stop sending `accel` messages.

```json
{
  "type": "stop-stream"
}
```

#### [C] Type: `accel`

Nexus should send this message for streaming acceleration data.

```json
{
  "type": "accel",
  "titanID": "6513270e-269e-0d37-f2a7-4de452e6b438",
  "magnitude": 0.63167553298292
}
```

#### [S] Type: `arm`

Server will send this message to make Nexus armed. Nexus should then send `movement-trigger` messages in case the acceleration goes higher than sensitivity.

```json
{
  "type": "arm"
}
```

#### [S] Type: `disarm`

Server will send this message to make Nexus disarmed, and stop the ongoing alarm.

```json
{
  "type": "disarm"
}
```

#### [C] Type: `movement-trigger`

Nexus should send this message when it is armed, and acceleration of one sensor goes above sensitivity.

```json
{
  "type": "movement-trigger",
  "titanID": "6513270e-269e-0d37-f2a7-4de452e6b438",
  "magnitude": 7.63167553298292
}
```

### User side

Endpoint: `wss://octoprotect.starcatmeow.cn/ws/user`

Authentication: Concat the username and password using `:`, then Base64-encode it, and put in Authorization header, with prefix `Basic `

#### [C] Type: `pair`

User will send this message to pair with a specific device. Server should reply with `pair-success` if succeeds.

```json
{
  "type": "pair",
  "nexusMac": "01:23:45:67:89:ab",
  "pairSecret": "FEC48C7C-F4B3-47C6-85B2-95E9E4ABAEB1",
  "nickName": "Catmeow's Nexus"
}
```

#### [S] Type: `pair-success`

Server should send this message to indicate pairing succeeds.

```json
{
  "type": "pair-success"
}
```

#### [C] Type: `unpair`

User will send this message to unpair with a specific device. Server should reply with `unpair-success` if succeeds.

```json
{
  "type": "unpair",
  "nexusID": 2
}
```

#### [S] Type: `unpair-success`

Server should send this message to indicate that the nexus was removed.

```json
{
  "type": "unpair-success",
  "nexusID": 2
}
```

#### [C] Type: `fetch-device-list`

User will send this message to fetch the list of paired devices. Server should reply with `device-list` if succeeds.

```json
{
  "type": "pair",
  "nexusMac": "01:23:45:67:89:ab",
  "pairSecret": "FEC48C7C-F4B3-47C6-85B2-95E9E4ABAEB1"
}
```

#### [S] Type: `device-list`

Response of `fetch-device-list`.

```json
{
  "type": "device-list",
  "data": [
    {
      "id": 1,
      "macAddress": "01:23:45:67:89:ae",
      "config": {
        "sensitivity": 1,
        "titanW": [ // all paired Titan W
          {
            "uuid": "6513270e-269e-0d37-f2a7-4de452e6b438",
            "enabled": true
          }
        ]
      },
      "online": false
    }
  ]
}
```

#### [C] Type: `update-config`

User will send this message to update config of a paired device. Server should respond with `update-config-success` if succeeds.

```json
{
  "type": "update-config",
  "nexusID": 2,
  "sensitivity": 1.3,
  "titanW": [
    {
      "uuid": "6513270e-269e-0d37-f2a7-4de452e6b438",
      "enabled": true
    }
  ]
}
```

#### [S] Type: `update-config-success`

Server should send this message to indicate the success of `update-config`.

```json
{
  "type": "update-config-success"
}
```

#### [C] Type: `request-state`

User will send this message to request state of a paired Nexus. Server should respond with `nexus-state`.

```json
{
  "type": "request-state",
  "nexusID": 1
}
```

#### [S] Type: `nexus-state`

Server should send this message as response of `request-state`.

```json
{
  "type": "nexus-state",
  "nexusID": 1,
  "titan": [
    {
      "id": "e3e70682-c209-4cac-629f-6fbed82c07cd",
      "name": "Titan A - Attached Accelerometer",
      "isWorking": true
    },
    {
      "id": "cd613e30-d8f1-6adf-91b7-584a2265b1f5",
      "name": "Titan A - Attached Accelerometer",
      "isWorking": true
    },
    {
      "id": "6513270e-269e-0d37-f2a7-4de452e6b438",
      "name": "Titan A - Attached Accelerometer",
      "isWorking": true
    }
  ],
  "isArmed": false
}
```

#### [S] Type: `conn-state`

Server should send this message when there is a Titan connection change.

```json
{
  "type": "conn-state",
  "nexusID": 1,
  "titanID": "6513270e-269e-0d37-f2a7-4de452e6b438",
  "isConnected": false
}
```

#### [C] Type: `start-stream`

User will send this message to enable Nexus acceleration data streaming. Server should then start sending `accel` messages.

```json
{
  "type": "start-stream",
  "nexusID": 1
}
```

#### [C] Type: `stop-stream`

User will send this message to disable Nexus acceleration data streaming. Server should then stop sending `accel` messages.

```json
{
  "type": "stop-stream",
  "nexusID": 1
}
```

#### [S] Type: `accel`

Server should send this message for streaming acceleration data.

```json
{
  "type": "accel",
  "nexusID": 1,
  "titanID": "6513270e-269e-0d37-f2a7-4de452e6b438",
  "magnitude": 0.63167553298292
}
```

#### [C] Type: `arm`

User will send this message to make a paired Nexus armed. Server should then send `movement-trigger` messages in case the acceleration goes higher than sensitivity.

```json
{
  "type": "arm",
  "nexusID": 1
}
```

#### [S] Type: `disarm`

User will send this message to make a paired Nexus disarmed, and stop the ongoing alarm.

```json
{
  "type": "disarm",
  "nexusID": 1
}
```

#### [S] Type: `movement-trigger`

Server should send this message when a paired Nexus is armed, and acceleration of one sensor goes above sensitivity.

```json
{
  "type": "movement-trigger",
  "nexusID": 1,
  "titanID": "6513270e-269e-0d37-f2a7-4de452e6b438",
  "magnitude": 7.63167553298292
}
```

## HTTP Admin Endpoint

Base URL: `https://octoprotect.starcatmeow.cn`

Authentication: Put admin token in Authorization header, with prefix `Token `.

### Provision Nexus (`POST /nexus`)

Admin can provision a new Nexus (generate pair secret, and enroll in the DB), with JSON body

```json
{
  "nexusMac": "01:23:45:67:89:ae"
}
```

The response will be like

```json
{
    "ID": 1,
    "CreatedAt": "2023-11-24T00:00:05.900115111Z",
    "UpdatedAt": "2023-11-24T00:00:05.900115111Z",
    "DeletedAt": null,
    "macAddress": "01:23:45:67:89:ae",
    "pairSecret": "f8be6447-90e3-4207-9aaa-bb5d5b90b990",
    "users": [],
    "config": {
        "sensitivity": 1,
        "titanW": []
    }
}
```

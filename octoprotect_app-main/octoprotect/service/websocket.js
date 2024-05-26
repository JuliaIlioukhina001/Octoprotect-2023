import { BACKEND_URL } from "../config"
import { encode } from 'base-64';
import { createAction } from "@reduxjs/toolkit";
import Toast from "react-native-root-toast";

export const wsConnectedAction = createAction('server/connected');
export const wsDisconnectedAction = createAction('server/disconnected');

export const connectAction = createAction('client/connect');
export const disconnectAction = createAction('client/disconnect');

export const errorAction = createAction('server/error');

export const pairAction = createAction('client/pair');
export const pairSuccessAction = createAction('server/pair-success');

export const fetchDeviceListAction = createAction('client/fetch-device-list');
export const deviceListAction = createAction('server/device-list');

export const updateConfigAction = createAction('client/update-config');
export const updateConfigSuccessAction = createAction('server/update-config-success');

export const requestStateAction = createAction('client/request-state');
export const nexusStateAction = createAction('server/nexus-state');

export const connStateAction = createAction('server/conn-state');

export const startStreamAction = createAction('client/start-stream');
export const stopStreamAction = createAction('client/stop-stream');
export const accelAction = createAction('server/accel');

export const armAction = createAction('client/arm');
export const disarmAction = createAction('client/disarm');
export const movementTriggerAction = createAction('server/movement-trigger');


export class Socket {
  constructor(){
    this.socket = null
    this.shouldConnect = false
  }
  connect(cred, openHandler, closeHandler, messageHandler){
    const base64Cred = encode(cred.username+':'+cred.password)
    this.shouldConnect = true
    this.socket = new WebSocket(BACKEND_URL, BACKEND_URL.split(':')[0], {
      headers: {
        'Authorization': 'Basic '+base64Cred
      }
    })
    this.socket.onopen = openHandler
    this.socket.onclose = function () {
      Toast.show('Disconnected, reconnecting...')
      closeHandler()
      // auto reconnect
      if (this.shouldConnect) {
        setTimeout(() => {
          this.connect(cred, openHandler, closeHandler, messageHandler)
        }, 1000)
      }
    }
    this.socket.onmessage = e => {
      const data = JSON.parse(e.data)
      console.log("Received:", data)
      messageHandler(data)
    }
    this.socket.onerror = e => {
      Toast.show('Network error, reconnecting. Error:' + e.message)
      console.error(e.message)
      // auto reconnect
      if (this.shouldConnect) {
        setTimeout(() => {
          this.connect(cred, openHandler, closeHandler, messageHandler)
        }, 1000)
      }
    }
  }
  disconnect(){
    this.shouldConnect = false
    this.socket.disconnect()
  }
  send(message){
    this.socket.send(JSON.stringify(message))
    console.log("Sent:",message)
  }
}
export const WebsocketMiddleware = (socket) => (params) => (next) => (action) => {
  const { dispatch, getState } = params
  const { type, payload } = action
  const openHandler = () => {
    dispatch(wsConnectedAction())
  }
  const closeHandler = () => {
    dispatch(wsDisconnectedAction())
  }
  const messageHandler = data => {
    dispatch({
      type: 'server/' + data.type,
      payload: data
    })
    if (data.type === "conn-state" && getState().device.configDirty && data.isConnected) {
      dispatch(requestStateAction({nexusID: data.nexusID}))
    }
  }
  switch (type) {
    case connectAction().type:
      console.log(type, payload)
      socket.connect(
        payload,
        openHandler,
        closeHandler,
        messageHandler
      )
      break
    case disconnectAction().type:
      socket.disconnect()
      break
    default:
      if (type.startsWith("client/")) {
        socket.send({
          type: type.substring('client/'.length),
          ...payload
        })
      }
  }
  return next(action)
}
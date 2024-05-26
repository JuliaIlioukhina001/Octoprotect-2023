import { createSlice } from "@reduxjs/toolkit";
import { accelAction, connStateAction, disarmAction, movementTriggerAction, nexusStateAction, updateConfigAction } from "../service/websocket";

export const deviceSlice = createSlice({
  name: 'device',
  initialState: {
    device: null,
    nexusState: null,
    acceleration: {},
    configDirty: false,
  },
  reducers: {
    setCurrentDevice: (state, {payload}) => {
      state.device = payload.device;
    },
    pairTitanW: (state, {payload}) => {
      console.log(state, payload)
      if (!state.device.config.titanW.find(t => t.uuid === payload.uuid)) {
        state.device.config.titanW.push({
          uuid: payload.uuid,
          nickName: payload.nickName,
          enabled: false
        })
      }
    },
    toggleTitanW: (state, {payload}) => {
      state.device.config.titanW.find(t=>t.uuid === payload.uuid).enabled = payload.enabled
    },
    deleteTitanW: (state, {payload}) => {
      state.device.config.titanW = state.device.config.titanW.filter(t => t.uuid !== payload.uuid)
    }
  },
  extraReducers: builder => {
    builder
      .addCase(nexusStateAction, (state, {payload}) => {
        state.nexusState = payload
        state.configDirty = false
      })
      .addCase(movementTriggerAction, (state, {payload}) => {
        state.nexusState.isTriggered = true
      })
      .addCase(connStateAction, (state, {payload}) => {
        const titan = (state.nexusState?.titan || []).find(titan => titan.id === payload.titanID)
        console.log(state, payload, titan)
        if (!titan) return
        titan.isWorking = payload.isConnected
      })
      .addCase(accelAction, (state, {payload}) => {
        state.acceleration[payload.titanID] = payload.magnitude
      })
      .addCase(updateConfigAction, (state, {payload}) => {
        state.configDirty = true
      })
  }
})

import { createSlice } from "@reduxjs/toolkit";
import { connectAction, errorAction, movementTriggerAction, wsConnectedAction, wsDisconnectedAction } from "../service/websocket";

export const appSlice = createSlice({
  name: 'app',
  initialState: {
    connected: false,
    authed: false,
    loggingIn: false,
    showAlert: false,
    error: "",
  },
  reducers: {
    hideAlert: (state) => {
      state.showAlert = false
    }
  },
  extraReducers: builder => {
    builder
      .addCase(wsConnectedAction, state => {
        state.authed = true;
        state.connected = true;
        state.loggingIn = false;
      })
      .addCase(wsDisconnectedAction, state => {
        state.connected = false;
      })
      .addCase(connectAction, state => {
        state.loggingIn = true;
      })
      .addCase(errorAction, (state, {payload}) => {
        state.error = payload.message
      })
      .addCase(movementTriggerAction, (state) => {
        state.showAlert = true;
      });
  }
})
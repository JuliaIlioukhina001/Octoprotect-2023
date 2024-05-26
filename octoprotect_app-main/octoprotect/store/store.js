import {
  combineReducers,
  configureStore,
  getDefaultMiddleware,
} from "@reduxjs/toolkit";
import { deviceListSlice } from "./deviceListSlice";
import { appSlice } from "./appSlice";
import { Socket, WebsocketMiddleware } from "../service/websocket";
import { deviceSlice } from "./deviceSlice";
import { devicePairSlice } from "./devicePairSlice";

export const store = configureStore({
  reducer: {
    app: appSlice.reducer,
    deviceList: deviceListSlice.reducer,
        device: deviceSlice.reducer,
    devicePair: devicePairSlice.reducer,
  },
  middleware: (getDefaultMiddleware) => [
    new WebsocketMiddleware(new Socket()),
    ...getDefaultMiddleware(),
  ],
});

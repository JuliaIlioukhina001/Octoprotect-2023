import { createSlice } from "@reduxjs/toolkit";
import { pairSuccessAction, pairAction } from "../service/websocket";

export const devicePairSlice = createSlice({
  name: "devicePair",
  initialState: {
    mode: "",
    nexusMac: "",
    pairSecret: "",
    uuid: "",
    pairSuccess: false,
  },
  reducers: {
    qrCodeScanned: (state, { payload }) => {
      state.mode = payload.mode;
      if (payload.mode === "nexus") {
        state.nexusMac = payload.macAddress;
        state.pairSecret = payload.pairSecret;
      } else if (payload.mode === "titanw") {
        state.uuid = payload.uuid
      }
    },
    resetPairSuccess: (state) => {
      state.pairSuccess = false;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(pairSuccessAction, (state, data) => {
      state.pairSuccess = true;
    });
  },
});

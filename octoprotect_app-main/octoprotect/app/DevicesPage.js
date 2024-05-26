import { StyleSheet, Text, View, TouchableOpacity } from "react-native";
import { useEffect, useState } from "react";
import Toggle from "../elements/Toggle";
import { useDispatch, useSelector } from "react-redux";
import { fetchDeviceListAction } from "../service/websocket";
import AddNewDeviceButton from "../elements/AddNewDeviceButton";
import { deviceSlice } from "../store/deviceSlice";
import AlertOverlay from "./AlertOverlay.js";

const selectDevices = (state) => state.deviceList.devices;
const DevicesPage = ({ navigation }) => {
  const devices = useSelector(selectDevices);
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchDeviceListAction());
  }, []);
  const handleClick = (data) => {
    dispatch(deviceSlice.actions.setCurrentDevice({ device: data }));
    navigation.navigate("Device");
  };

  return (
    <View style={styles.deviceList}>
      {devices.map((device) => {
        return (
          <View style={styles.deviceContainer} key={device.id}>
            <TouchableOpacity onPress={() => handleClick(device)}>
              {/* {<DeviceSpecificPage deviceData={selectedDevice} />} */}
              <Text style={styles.header2}>{device.nickName}</Text>
            </TouchableOpacity>
          </View>
        );
      })}
      <AddNewDeviceButton
        onPress={() => {
          navigation.navigate("Scan QR Code");
        }}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  title: {
    flex: 1,
    justifyContent: "center",
    padding: 40,
    color: "black",
  },
  header2: {
    fontSize: 20,
    marginBottom: 20,
    fontWeight: "300",
  },
  deviceList: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    width: "100%",
  },
  deviceContainer: {
    borderRadius: 8,
    paddingVertical: 15,
    paddingHorizontal: 25,
    width: "90%",
    marginVertical: 10,
    marginLeft: 10,
    marginRight: 10,
    backgroundColor: "#DDEFE5",
  },
  toggleStyles: {
    marginTop: "30%",
    width: "50%",
    paddingTop: "-20%",
  },
});

export default DevicesPage;

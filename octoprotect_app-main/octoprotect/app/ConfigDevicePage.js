import { useCallback, useEffect, useMemo, useState } from "react";
import { StyleSheet, View } from "react-native";
import { Button, Chip, Divider, Icon, Surface, Switch, Text, Appbar, ProgressBar, MD3Colors } from "react-native-paper";
import { useDispatch, useSelector } from "react-redux";
import { armAction, disarmAction, requestStateAction, startStreamAction, stopStreamAction, updateConfigAction } from "../service/websocket";
import { BlinkText } from "../components/BlinkText";
import Slider from "@react-native-community/slider";
import { deviceSlice } from "../store/deviceSlice";
import { AccelerationDisplay } from "../components/AccelerationDisplay";

const deviceSelector = state => state.device.device
const accelerationSelector = state => state.device.acceleration
const nexusStateSelector = state => state.device.nexusState
const ConfigDevicePage = ({ navigation }) => {
  const device = useSelector(deviceSelector)
  const nexusState = useSelector(nexusStateSelector)
  const acceleration = useSelector(accelerationSelector)
  const [sensitivity, setSensitivity] = useState(0)
  useEffect(() => {
    setSensitivity(device.config.sensitivity)
  }, [device.config.sensitivity])
  useEffect(() => {
    dispatch(startStreamAction({nexusID: device.id}))
    return () => {
      dispatch(stopStreamAction({nexusID: device.id}))
    }
  }, [])
  const titans = useMemo(() => {
    const result = {}
    if (nexusState?.titan) {
        nexusState.titan.forEach(titan => {
        result[titan.id] = {
          name: titan.name,
          isConnected: titan.isWorking,
        }
      })
    }
    device.config.titanW.forEach(titanW => {
      if (!result[titanW.uuid]) {
        result[titanW.uuid] = {}
      }
      result[titanW.uuid] = {
        ...result[titanW.uuid],
        name: titanW.nickName,
        enabled: titanW.enabled,
      }
    })
    return Object.entries(result).map(([id, titan]) => ({
      ...titan,
      id,
    }))
  }, [device, nexusState]);
  const dispatch = useDispatch()
  const applyConfig = useCallback(() => {
    dispatch(updateConfigAction({
      ...device.config,
      nexusID: device.id
    }))
    navigation.goBack()
  },[dispatch, device])
  return (
    <View style={styles.container}>
      <View style={styles.deviceNameContainer}>
      <Text style={styles.deviceName}>{device.nickName}</Text>
      <Chip>{device.online ? 'Connected' : 'Disconnected'}</Chip>
      </View>
      <Text style={styles.extraInfo}>MAC: {device.macAddress}</Text>
      {nexusState && <Surface mode="flat" style={styles.armedStatusContainer}>
        <View style={styles.flexRow}>
          <Text>Sensitivity: {sensitivity.toFixed(2)}</Text>
          <Slider style={{flexGrow: 1}} minimumValue={0} maximumValue={2} value={sensitivity} onValueChange={setSensitivity}/>
        </View>
      </Surface>}
      <Divider />
      <View style={styles.flexRow}>
        <Text variant="titleLarge" style={{flexGrow: 1}}>Titans</Text>
        <Button onPress={() => navigation.navigate("Scan QR Code")}>Pair Titan W</Button>
      </View>
      {/* {nexusState && nexusState.titan.map(titan => <Surface mode="flat" style={styles.titanContainer} key={titan.id}>
        <View style={styles.flexRow}><Icon source={titan.isWorking ? "check" : "close"}/><Text variant="bodyLarge">{titan.name}</Text></View>
        <Text variant="labelSmall">ID: {titan.id}</Text>
        {acceleration[titan.id] !== undefined && <View>
          <Text variant="bodySmall">Acceleration: {acceleration[titan.id].toFixed(2)}</Text>
          <ProgressBar progress={acceleration[titan.id]/2} color={acceleration[titan.id] >= sensitivity ? MD3Colors.error50 : MD3Colors.primary50} />
        </View>}
        
      </Surface>)} */}
      {titans.map(titan => <Surface mode="flat" style={styles.titanContainer} key={titan.id}>
        <View style={styles.flexRow}>
          <Icon source={titan.isConnected ? "check" : "close"}/>
          <Text variant="bodyLarge">{titan.name}</Text>
        </View>
        <Text variant="labelSmall">ID: {titan.id}</Text>
        <AccelerationDisplay titanID={titan.id} sensitivity={sensitivity} />
        {titan.enabled !== undefined && <View style={styles.flexRow}>
          <Text>{titan.enabled ? "Enabled": "Disabled"}</Text>
          <Switch
            value={titan.enabled}
            onValueChange={value => dispatch(deviceSlice.actions.toggleTitanW({uuid: titan.id, enabled: value}))}
            />
          <View style={{flexGrow: 1}} />
          <Button
            textColor={MD3Colors.error50}
            onPress={() => dispatch(deviceSlice.actions.deleteTitanW({uuid: titan.id}))}
          >
            Remove
          </Button>
        </View>}
      </Surface>)}
      <Button mode="outlined" icon="check" onPress={applyConfig}>Apply Changes</Button>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: 20,
    // backgroundColor: 'lightblue',
    width: '100%',
    display: "flex",
    gap: 8
  },
  deviceNameContainer: {
    display: "flex",
    flexDirection: "row",
  },
  deviceName: {
    fontSize: 24,
    fontWeight: "bold",
    flexGrow: 1
  },
  header2: {
    fontSize: 20,
    marginBottom: 20,
  },
  header1: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 20,
    paddingTop: 60,
    paddingLeft: 40,
  },
  armedStatusContainer: {
    padding: 16,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    borderRadius: 16
  },
  titanContainer: {
    padding: 16,
    display: "flex",
    flexDirection: "column",
    borderRadius: 16
  },
  flexRow: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    gap: 4
  },
  squareButton: {
    borderRadius: 8,
  },
  squareButtonText: {
    marginHorizontal: 12,
    marginVertical: 5,
  },
  deviceContainer: {
    borderRadius: 8,
    paddingVertical: 15,
    paddingHorizontal: 25,
    width: '90%',
    marginVertical: 10,
    marginLeft: 10,
    marginRight: 10,
    backgroundColor: '#DDEFE5',
    // shadowColor: '#171717',
    // shadowOffset: {width: -2, height: 4},
    // shadowOpacity: 0.9,
    // shadowRadius: 3,
  },
  toggleStyles: {
    marginTop: '30%',
    width: '50%',
    paddingTop: '-20%',
  },
});

export default ConfigDevicePage;

import { useCallback, useEffect } from "react";
import { StyleSheet, View } from "react-native";
import {
  Button,
  Chip,
  Divider,
  Icon,
  Surface,
  Switch,
  Text,
  Appbar,
} from "react-native-paper";
import { useDispatch, useSelector } from "react-redux";
import {
  armAction,
  disarmAction,
  requestStateAction,
  updateConfigAction,
} from "../service/websocket";
import { BlinkText } from "../components/BlinkText";

const deviceSelector = state => state.device.device
const nexusStateSelector = state => state.device.nexusState
const configDirtySelector = state => state.device.configDirty
const DeviceSpecificPage = ({ navigation }) => {
  const device = useSelector(deviceSelector)
  const nexusState = useSelector(nexusStateSelector)
  const configDirty = useSelector(configDirtySelector)
  const dispatch = useDispatch()
  useEffect(() => {
    navigation.setOptions({
      actions: [
        <Appbar.Action
          icon="pencil"
          key="config"
          onPress={() => {
            navigation.navigate("Config");
          }}
        />,
        <Appbar.Action
          icon="power"
          key="restart"
          onPress={() => {
            dispatch(
              updateConfigAction({
                ...device.config,
                nexusID: device.id,
              })
            );
          }}
        />,
      ],
    });
  }, [dispatch]);
  useEffect(() => {
    dispatch(requestStateAction({nexusID: device.id}))
  }, [device.id])
  const changeArmStatus = useCallback((newState) => {
    dispatch((newState ? armAction : disarmAction)({nexusID: device.id}))
    dispatch(requestStateAction({nexusID: device.id}))
  }, [device.id])
  useEffect(() => {
    if (configDirty && nexusState.titan.filter(t => !t.isWorking) === 0) {
      dispatch(requestStateAction({nexusID: device.id}))
    }
  }, [nexusState])
  return (
    <View style={styles.container}>
      <View style={styles.deviceNameContainer}>
      <BlinkText style={styles.deviceName} blink={nexusState?.isTriggered}>Mariya's Items</BlinkText>
      <Chip>{device.online ? 'Connected' : 'Disconnected'}</Chip>
      </View>
      <Text style={styles.extraInfo}>MAC: {device.macAddress}</Text>
      {nexusState && (
        <Surface mode="flat" style={styles.armedStatusContainer}>
          <View style={styles.flexRow}>
            <Text style={{ flexGrow: 1 }}>
              {nexusState.isArmed ? "Armed" : "Disarmed"}
            </Text>
            <Switch
              value={nexusState.isArmed}
              onValueChange={changeArmStatus}
            />
          </View>
        </Surface>
      )}
      {nexusState && (
        <>
          <Divider />
          <Text variant="titleLarge">Connected Titan</Text>

          {nexusState.titan.map((titan) => (
            <Surface mode="flat" style={styles.titanContainer} key={titan.id}>
              <View style={styles.flexRow}>
                <Icon source={titan.isWorking ? "check" : "close"} />
                <Text variant="bodyLarge">{titan.name}</Text>
              </View>
              <Text variant="labelSmall">ID: {titan.id}</Text>
            </Surface>
          ))}
        </>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: 20,
    // backgroundColor: 'lightblue',
    width: "100%",
    display: "flex",
    gap: 8,
  },
  deviceNameContainer: {
    display: "flex",
    flexDirection: "row",
    paddingVertical: 20,
  },
  deviceName: {
    fontSize: 24,
    fontWeight: "bold",
    flexGrow: 1,
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
    borderRadius: 16,
  },
  titanContainer: {
    padding: 16,
    display: "flex",
    flexDirection: "column",
    borderRadius: 16,
  },
  flexRow: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    gap: 4,
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
    width: "90%",
    marginVertical: 10,
    marginLeft: 10,
    marginRight: 10,
    backgroundColor: "#DDEFE5",
    // shadowColor: '#171717',
    // shadowOffset: {width: -2, height: 4},
    // shadowOpacity: 0.9,
    // shadowRadius: 3,
  },
  toggleStyles: {
    marginTop: "30%",
    width: "50%",
    paddingTop: "-20%",
  },
});

export default DeviceSpecificPage;

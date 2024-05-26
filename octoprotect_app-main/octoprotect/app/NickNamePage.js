import { StyleSheet, Text, View, TextInput, Button } from "react-native";
import { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchDeviceListAction, pairAction } from "../service/websocket";
import { devicePairSlice } from "../store/devicePairSlice";
import { deviceSlice } from "../store/deviceSlice";

const pairData = (state) => state.devicePair;
const NickNamePage = ({ navigation }) => {
  const [name, setName] = useState("");
  const deviceData = useSelector(pairData);
  const dispatch = useDispatch();
  useEffect(() => {
    if (deviceData && deviceData.pairSuccess) {
      alert(`Sucessfully paired!`);
      navigation.navigate("Devices");
      dispatch(devicePairSlice.actions.resetPairSuccess());
      dispatch(fetchDeviceListAction());
    }
  }, [deviceData]);

  const nickNameData = {
    nickName: name,
  };
  const onSubmit = () => {
    if (!nickNameData.nickName) return;
    if (deviceData.mode === "nexus") {
      alert(`Pair request sent`);
      console.log(deviceData);
      dispatch(
        pairAction({
          nickName: name,
          ...deviceData,
        })
      );
    } else if (deviceData.mode === "titanw") {
      dispatch(
        deviceSlice.actions.pairTitanW({
          uuid: deviceData.uuid,
          nickName: name
        })
      )
      navigation.goBack()
      navigation.goBack()
    }
  };
  return (
    <View style={styles.container}>
      <Text style={styles.header1}>Step 2:</Text>
      <Text style={styles.header2}>
        Give your device a nickname to help you identify it more easily:
      </Text>
      <TextInput
        style={styles.input}
        placeholder="Enter a nickname"
        value={name}
        onChangeText={(text) => setName(text)}
      />
      <Button title="Save" onPress={() => onSubmit()} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    padding: 40,
  },
  header1: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 20,
  },
  header2: {
    fontSize: 20,
    marginBottom: 20,
  },
  input: {
    borderWidth: 1,
    borderColor: "#ccc",
    padding: 10,
    marginBottom: 20,
    width: "100%",
    fontSize: 18,
  },
});

export default NickNamePage;

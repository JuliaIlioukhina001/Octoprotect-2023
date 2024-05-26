import { Pressable, StyleSheet, Text, View } from "react-native";

const AddNewDeviceButton = ({ onPress }) => {
  return (
    <Pressable
      onPress={() => onPress()}
      style={({ pressed }) => [
        {
          backgroundColor: pressed ? "#549B0D" : "#5EBC00",
          alignItems: "center",
        },
        styles.wrapperCustom,
      ]}
    >
      <Text style={styles.text}>Add a Device</Text>
    </Pressable>
  );
};

const styles = StyleSheet.create({
  text: {
    fontSize: 19,
    fontWeight: "500",
    color: "white",
  },
  wrapperCustom: {
    borderRadius: 20,
    padding: 12,
    width: 200,
    elevation: 5,
    shadowOffset: 20,
    marginTop: 5,
  },
});

export default AddNewDeviceButton;

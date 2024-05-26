import { useCallback, useState } from "react";
import { StyleSheet, View } from "react-native";
import { Button, TextInput } from "react-native-paper";
import { useDispatch, useSelector } from "react-redux";
import { connectAction } from "../service/websocket";

const selectLoggingIn = (state) => state.app.loggingIn;

const LoginPage = () => {
  const [username, setUsername] = useState('tesxuser2');
  const [password, setPassword] = useState('testpass');
  const loggingIn = useSelector(selectLoggingIn)
  const dispatch = useDispatch()
  const login = () => {
    console.log("login called");
    dispatch(
      connectAction({
        username,
        password,
      })
    );
  };
  return (
    <View style={styles.container}>
      <TextInput
        mode="outlined"
        label="Username"
        value={username}
        onChangeText={setUsername}
      />
      <TextInput
        mode="outlined"
        label="Password"
        value={password}
        onChangeText={setPassword}
      />
      <Button mode="outlined" loading={loggingIn} onPress={login}>
        Login
      </Button>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    display: "flex",
    gap: 16,
    marginHorizontal: 16,
  },
});

export default LoginPage;

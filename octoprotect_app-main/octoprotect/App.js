import { NavigationContainer } from "@react-navigation/native";
import { MD3LightTheme, PaperProvider, Snackbar } from "react-native-paper";
import { Provider, useSelector } from "react-redux";
import { store } from "./store/store";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import LoginPage from "./app/LoginPage";
import DevicesPage from "./app/DevicesPage";
import NavBar from "./components/navBar";
import QRScanner from "./app/QRScanner";
import NickNamePage from "./app/NickNamePage";
import "react-native-url-polyfill/auto";
import DeviceSpecificPage from "./app/DeviceSpecificPage";
import ConfigDevicePage from "./app/ConfigDevicePage";
import AlertOverlay from "./app/AlertOverlay";
import { ErrorSnackbar } from "./components/ErrorSnackbar"

export default function App() {
  return (
    <Provider store={store}>
      <PaperProvider theme={MD3LightTheme}>
        <NavigationContainer theme={MD3LightTheme}>
          <Router />
          <ErrorSnackbar />
          <AlertOverlay />
        </NavigationContainer>
      </PaperProvider>
    </Provider>
  );
}

const Stack = createNativeStackNavigator();
const selectAuthed = (state) => state.app.authed;
const selectError = (state) => state.app.error;
export const Router = () => {
  const authed = useSelector(selectAuthed);
  return (<>
    <Stack.Navigator
      screenOptions={{
        header: (props) => <NavBar {...props} />,
      }}
    >
      {!authed && (
        <>
          <Stack.Screen name="Login" component={LoginPage} />
        </>
      )}
      {authed && (
        <>
          <Stack.Screen name="Devices" component={DevicesPage} />
          <Stack.Screen name="Device" component={DeviceSpecificPage} />
          <Stack.Screen name="Config" component={ConfigDevicePage} />
        </>
      )}
      <Stack.Screen name="Scan QR Code" component={QRScanner} />
      <Stack.Screen name="Nickname Page" component={NickNamePage} />
    </Stack.Navigator>
  </>);
};

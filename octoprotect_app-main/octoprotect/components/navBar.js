import { Appbar, Text } from 'react-native-paper';
import { getHeaderTitle } from '@react-navigation/elements';
import { useSelector } from 'react-redux';
import { StyleSheet, View } from 'react-native';

const selectConnected = state => state.app.connected
const selectAuthed = state => state.app.authed

export const NavBar = ({ navigation, route, options, back }) => {
  const title = getHeaderTitle(options, route.name);
  const connected = useSelector(selectConnected)
  const authed = useSelector(selectAuthed)
  return (
    <Appbar.Header>
      {back ? <Appbar.BackAction onPress={navigation.goBack} /> : null}
      <View
        style={[
          StyleSheet.absoluteFill,
        ]}
        pointerEvents="box-none"
      >
        <Appbar.Content
          title={(authed && !connected) ? 'Disconnected' : title}
          style={{
            alignItems: "center",
            justifyContent: "center",
          }}
        />
      </View>
      <View style={{flexGrow: 1}} />
      {options.actions}
    </Appbar.Header>
  );
}

const styles = StyleSheet.create({
  actions: {
    position: 'absolute',
    right: 0,
    top: 0
  }
})

export default NavBar
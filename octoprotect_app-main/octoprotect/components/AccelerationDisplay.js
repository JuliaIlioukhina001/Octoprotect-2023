import { useSelector } from "react-redux"
import { View, Text } from "react-native"
import { MD3Colors, ProgressBar } from "react-native-paper"

const accelerationSelector = state => state.device.acceleration
export const AccelerationDisplay = (props) => {
  const acceleration = useSelector(accelerationSelector)
  return <>
    {acceleration[props.titanID] !== undefined && <View>
      <Text variant="bodySmall">Acceleration: {acceleration[props.titanID].toFixed(2)}</Text>
      <ProgressBar progress={acceleration[props.titanID]/2} color={acceleration[props.titanID] >= props.sensitivity ? MD3Colors.error50 : MD3Colors.primary50} />
    </View>}
  </>
}
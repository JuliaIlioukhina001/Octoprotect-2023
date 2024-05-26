import { useEffect } from 'react';
import { Text } from 'react-native-paper';
import Animated, {useSharedValue, useAnimatedStyle, withRepeat, withTiming, Easing} from 'react-native-reanimated';
export const BlinkText = ({blink, ...props}) => {
  const offset = useSharedValue(1);
  const animatedStyles = useAnimatedStyle(() => ({
    opacity: offset.value
  }))
  useEffect(() => {
    offset.value = withRepeat(
      withTiming(0, {duration: 500, easing: Easing.out(Easing.cubic)}),
      -0.01,
      true
    )
  }, [])
  return (<>
    {!blink && <Text {...props} />}
    {blink && <Animated.Text {...props} style={[
      props.style,
      {color: '#ff0000'},
      animatedStyles,
    ]}/>}
  </>)
}
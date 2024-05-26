import { useEffect } from "react"
import { useDispatch, useSelector } from "react-redux"
import { errorAction } from "../service/websocket"
import { Snackbar } from "react-native-paper"

const selectError = state => state.app.error
export const ErrorSnackbar = () => {
  const error = useSelector(selectError)
  const dispatch = useDispatch()
  useEffect(() => {
    if (error !== '')
      setTimeout(() => dispatch(errorAction({message: ''})), 3000)
  }, [error])
  return <Snackbar visible={error && error.length}>{error}</Snackbar>
}
import { View, Switch } from "react-native";
import React, { useState } from "react";
const Toggle = () => {
  const [status, setToggleStatus] = useState(false);
  const toggleStatus = () => setToggleStatus((status) => !status);
  return (
    <View>
      <Switch
        trackColor={"D9D9D9"}
        thumbColor={status ? "5EBC00" : "858585"}
        onValueChange={toggleStatus}
        value={status}
      />
    </View>
  );
};

export default Toggle;

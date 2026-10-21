import { Text, View, TextInput, Button, StyleSheet } from 'react-native';
import { type JSX, useState } from 'react';
import SampleTurboModule from '../specs/NativeChatModule';

const Chat = () => {
  const [value, setValue] = useState('');
  const [reversedValue, setReversedValue] = useState('');

  const onPress = () => {
    const revString = SampleTurboModule.reverseString(value);
    setReversedValue(revString);
  };
  return (
    <View
      style={{
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
      }}>
      <Text>No waylo</Text>
      <Text>Write down here the text you want to reverse</Text>
      <TextInput
        style={styles.textInput}
        placeholder="Write your text here"
        onChangeText={setValue}
        value={value}
      />
      <Button title="Reverse" onPress={onPress} />
      <Text>Reversed text: {reversedValue}</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    fontSize: 18,
    marginBottom: 20,
  },
  textInput: {
    borderColor: 'black',
    borderWidth: 1,
    borderRadius: 5,
    padding: 10,
    marginTop: 10,
  },
});

export default Chat;

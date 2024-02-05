import {
  HStack,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  NumberIncrementStepper,
  NumberDecrementStepper,
} from "@chakra-ui/react";
import { TimeStamps } from "../../../models/time";

interface TimeInputProps {
  timeStamp: TimeStamps;
  setTimeStamp: (time: number, action: string) => void;
}

function TimeInput({ timeStamp, setTimeStamp }: TimeInputProps) {
  return (
    <HStack shouldWrapChildren mb={2}>
      <NumberInput
        bg="white"
        size="xs"
        maxW={16}
        min={0}
        max={12}
        defaultValue={timeStamp.hours}
        onChange={(valueAsString) =>
          setTimeStamp(Number(valueAsString), "hours")
        }
      >
        <NumberInputField />
        <NumberInputStepper>
          <NumberIncrementStepper />
          <NumberDecrementStepper />
        </NumberInputStepper>
      </NumberInput>
      <NumberInput
        size="xs"
        maxW={16}
        min={0}
        max={60}
        defaultValue={timeStamp.minutes}
        onChange={(valueAsString) =>
          setTimeStamp(Number(valueAsString), "minutes")
        }
        bg="white"
      >
        <NumberInputField />
        <NumberInputStepper>
          <NumberIncrementStepper />
          <NumberDecrementStepper />
        </NumberInputStepper>
      </NumberInput>
      <NumberInput
        size="xs"
        maxW={16}
        min={0}
        max={60}
        defaultValue={timeStamp.seconds}
        onChange={(valueAsString) =>
          setTimeStamp(Number(valueAsString), "seconds")
        }
        bg="white"
      >
        <NumberInputField />
        <NumberInputStepper>
          <NumberIncrementStepper />
          <NumberDecrementStepper />
        </NumberInputStepper>
      </NumberInput>
    </HStack>
  );
}

export default TimeInput;

import {
  Stat,
  StatLabel,
  StatNumber,
  StatHelpText,
  VStack,
} from "@chakra-ui/react";
import { ReactNode } from "react";

interface BaseStatProps {
  label: string;
  value: string | number;
  children: ReactNode;
}

function BaseStat({ label, value, children }: BaseStatProps) {
  return (
    <Stat>
      <VStack>
        <StatLabel>{label}</StatLabel>
        <StatNumber>{value}</StatNumber>
        <StatHelpText>{children}</StatHelpText>
      </VStack>
    </Stat>
  );
}

export default BaseStat;

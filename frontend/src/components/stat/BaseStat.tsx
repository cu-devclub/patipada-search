import {
  Stat,
  StatLabel,
  StatNumber,
  StatHelpText,
  Flex
} from "@chakra-ui/react";
import { ReactNode } from "react";

interface BaseStatProps {
  label: string;
  value: string | number;
  children: ReactNode;
}

function BaseStat({ label, value, children }: BaseStatProps) {
  return (
    <Flex h='full' w='full'>
      <Stat>
        <StatLabel fontSize="3xl" textAlign='center'>
          {label}
        </StatLabel>
        <StatNumber textAlign='center' color='black' fontSize="2xl" fontWeight="bold" p={4}>{value}</StatNumber>
        <StatHelpText color='black'>{children}</StatHelpText>
      </Stat>
    </Flex>
  )
}

export default BaseStat;

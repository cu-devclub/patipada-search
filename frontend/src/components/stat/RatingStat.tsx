import {
  Stat,
  VStack,
  StatLabel,
  StatNumber,
  StatHelpText,
} from "@chakra-ui/react";

interface RatingStatProps {
  label: string;
  value: number;
  helper?: string;
}

function RatingStat({ label, value, helper }: RatingStatProps) {
  const color = value >= 4 ? "green" : value >= 2 ? "orange" : "red";
  
  return (
    <Stat>
      <VStack>
        <StatLabel fontSize="xl" fontWeight={"bold"}>
          {label}
        </StatLabel>
        <StatNumber fontSize="xl" color={color}>
          {value}
        </StatNumber>
        <StatHelpText fontSize="xl">{helper}</StatHelpText>
      </VStack>
    </Stat>
  );
}

export default RatingStat;

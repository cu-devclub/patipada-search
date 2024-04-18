import {
  CircularProgress,
  CircularProgressLabel,
  StatNumber,
  StatHelpText,
  Stat,
  Flex,
  Center,
  Grid,
  GridItem,
  Heading,
  HStack,
} from "@chakra-ui/react";

interface RatingStatProps {
  label: string;
  value: number;
  helper?: string;
}

function RatingStat({ value, helper }: RatingStatProps) {
  const color = value >= 4 ? "green" : value >= 2 ? "orange" : "red";
  const valuePercentage = value * 20;

  return (
    <Flex bg="gray.450" shadow="xl" h="full" w="full" justify="center">
      <Grid templateRows="repeat(6, 1fr)" gap="2" w="full">
        <GridItem rowSpan={2}>
          <Center h="full">
            <Heading color="black" fontWeight="bold">
              คะแนนเฉลี่ย
            </Heading>
          </Center>
        </GridItem>
        <GridItem rowSpan={4} px={4}>
          <HStack gap={8} h="full">
            <CircularProgress
              value={valuePercentage}
              color={color}
              size="5xs"
              thickness="15"
            >
              <CircularProgressLabel fontSize="xl">
                {valuePercentage}%
              </CircularProgressLabel>
            </CircularProgress>
            <Center
              bg="white"
              shadow="xl"
              borderWidth="1px"
              borderColor="gray.300"
              h="65%"
            >
              <Stat textAlign="center">
                <StatNumber color="black" fontWeight="bold" p={4}>
                  {value}
                </StatNumber>
                <StatHelpText color="black">{helper}</StatHelpText>
              </Stat>
            </Center>
          </HStack>
        </GridItem>
      </Grid>
    </Flex>
  );
}

export default RatingStat;

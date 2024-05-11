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
  average: number;
  percentage: number;
  helper?: string;
}

function RatingStat({ average, percentage, helper }: RatingStatProps) {
  const color = average >= 4 ? "green" : average >= 2 ? "orange" : "red";
  // const valuePercentage = twoDecimal(average * 20);

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
          <HStack gap={8} h="full" justify='center'>
            <CircularProgress
              value={percentage}
              color={color}
              size="5xs"
              thickness="15"
            >
              <CircularProgressLabel fontSize="xl">
                {percentage}%
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
                  {/* {twoDecimal(value)} */}
                  {average}
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

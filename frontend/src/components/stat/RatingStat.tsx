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
} from "@chakra-ui/react";

interface RatingStatProps {
  label: string;
  value: number;
  helper?: string;
}

function RatingStat({ value, helper }: RatingStatProps) {
  const color = value >= 4 ? "green" : value >= 2 ? "orange" : "red";

  return (
    <Flex bg="#D9D9D9" shadow="xl" h="full" w="full" justify="center">
      <Grid templateRows="repeat(6, 1fr)" gap="2" w="full">
        <GridItem rowSpan={2}>
          <Center h="full">
            <Heading color ='black' fontWeight='bold'> คะแนนเฉลี่ย</Heading>
          </Center>
        </GridItem>
        <GridItem rowSpan={4}>
          <Flex direction="row" justify='center' align='center' gap='8%' h='full'>
            <CircularProgress value={value*20} color={color} size='4xs' thickness='15'>
              <CircularProgressLabel fontSize='4xl'>{value*20}%</CircularProgressLabel>
            </CircularProgress>
            <Center p='6%' bg='white' shadow='xl' borderWidth='1px' borderColor='gray.300' h='65%'>
               <Stat textAlign='center'>
                <StatNumber color='black' fontWeight='bold' p={4}>
                  {value}
                </StatNumber>
                <StatHelpText color='black'>{helper}</StatHelpText>
              </Stat>
             </Center>
          </Flex>
        </GridItem>
      </Grid>
    </Flex>
  );
}

export default RatingStat;

import { AuthSummary } from "../../models/user";
import {
  Stat,
  VStack,
  StatLabel,
  StatNumber,
  Flex,
  Circle,
  Spacer,
  Grid,
  GridItem,
} from "@chakra-ui/react";
interface UserstatProps {
  authSummary: AuthSummary;
}
const boxStyle = {
  p: "3",
  bg: "white",
  shadow: "xl",
  borderWidth: "1px",
  borderColor: "gray.300",
  width: "100%",
}

function Userstat({ authSummary }: UserstatProps) {
  return (
    <Grid 
      templateRows="repeat(9, 1fr)"
      templateColumns="1fr" 
      h='full'
      p={5}
      bg='#D9D9D9'>
      <GridItem rowSpan={5}>
        <Stat>
          <StatLabel fontSize="4xl" fontWeight="bold" color='black' textAlign='center' >Total User</StatLabel>
        </Stat>
        <Stat>
          <Flex justifyContent="center" alignItems="flex-start" display="flex">
            <Circle bg='#FFA902' size='210px' m='5' >
              <Circle bg='white' size='140px' fontSize="5xl" color='black' fontWeight={"bold"}>{authSummary.sumTotal}</Circle>
            </Circle>
          </Flex>
        </Stat>
      </GridItem>
      <GridItem rowSpan={4}>
        <Stat>
          <VStack spacing='6'>
            <Flex sx={boxStyle}>
              <StatLabel fontSize="2xl" color='black' >User</StatLabel>
                <Spacer />
              <StatNumber fontSize="2xl" color='black' >{authSummary.totalUser}</StatNumber>
            </Flex>
            <Flex sx={boxStyle}>
              <StatLabel fontSize="2xl" color='black' >Admin</StatLabel>
                <Spacer />
              <StatNumber fontSize="2xl" color='black' >{authSummary.totalAdmin}</StatNumber>
            </Flex>
            <Flex sx={boxStyle}>
              <StatLabel fontSize="2xl" color='black' >Super Admin</StatLabel>
                <Spacer />
              <StatNumber fontSize="2xl" color='black' >{authSummary.totalSuperAdmin}</StatNumber>
            </Flex>
          </VStack>
        </Stat>
      </GridItem>
    </Grid>
  );
}

export default Userstat;

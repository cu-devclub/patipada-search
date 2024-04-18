import { AuthSummary } from "../../models/user";
import {
  Stat,
  VStack,
  StatLabel,
  StatNumber,
  Flex,
  Circle,
  Spacer,
  Box
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
};

function Userstat({ authSummary }: UserstatProps) {
  return (
    <Flex direction={"column"} h="full" p={5} gap={4} bg="gray.450">
      <VStack>
        <Stat>
          <StatLabel
            fontSize="4xl"
            fontWeight="bold"
            color="black"
            textAlign="center"
          >
            Total User
          </StatLabel>
        </Stat>
        <Stat>
          <VStack>
            <Circle bg="#FFA902" size="4xs">
              <Circle
                bg="white"
                size="5xs"
                fontSize="5xl"
                color="black"
                fontWeight={"bold"}
              >
                {authSummary.sumTotal}
              </Circle>
            </Circle>
          </VStack>
        </Stat>
      </VStack>

      <Box>
        <Stat>
          <VStack spacing="4">
            <Flex sx={boxStyle}>
              <StatLabel fontSize="2xl" color="black">
                User
              </StatLabel>
              <Spacer />
              <StatNumber fontSize="2xl" color="black">
                {authSummary.totalUser}
              </StatNumber>
            </Flex>
            <Flex sx={boxStyle}>
              <StatLabel fontSize="2xl" color="black">
                Admin
              </StatLabel>
              <Spacer />
              <StatNumber fontSize="2xl" color="black">
                {authSummary.totalAdmin}
              </StatNumber>
            </Flex>
            <Flex sx={boxStyle}>
              <StatLabel fontSize="2xl" color="black">
                Super Admin
              </StatLabel>
              <Spacer />
              <StatNumber fontSize="2xl" color="black">
                {authSummary.totalSuperAdmin}
              </StatNumber>
            </Flex>
          </VStack>
        </Stat>
      </Box>
    </Flex>
  );
}

export default Userstat;

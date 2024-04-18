import { AuthSummary } from "../../models/user";
import BaseStat from "./BaseStat";
import { HStack, Text, VStack } from "@chakra-ui/react";
interface UserstatProps {
  authSummary: AuthSummary;
}

function Userstat({ authSummary }: UserstatProps) {
  return (
    <BaseStat label={"Total User"} value={authSummary.sumTotal}>
      <HStack>
        <VStack>
          <Text>User</Text>
          <Text>{authSummary.totalUser}</Text>
        </VStack>
        <VStack>
          <Text>Admin</Text>
          <Text>{authSummary.totalAdmin}</Text>
        </VStack>
        <VStack>
          <Text>Super Admin</Text>
          <Text>{authSummary.totalSuperAdmin}</Text>
        </VStack>
      </HStack>
    </BaseStat>
  );
}

export default Userstat;

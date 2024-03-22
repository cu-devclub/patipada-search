import { RequestSummary } from "../../models/request";
import BaseStat from "./BaseStat";
import { HStack, Text, VStack } from "@chakra-ui/react";
interface RequestStatProps {
  requestSummary: RequestSummary;
}
function RequestStat({ requestSummary }: RequestStatProps) {
  return (
    <BaseStat label={"Total Request"} value={requestSummary.requestAmount}>
      <HStack>
        <VStack>
          <Text>Pending</Text>
          <Text>{requestSummary.pendingAmount}</Text>
        </VStack>
        <VStack>
          <Text>Reviewed</Text>
          <Text>{requestSummary.reviewedAmount}</Text>
        </VStack>
      </HStack>
    </BaseStat>
  );
}

export default RequestStat;

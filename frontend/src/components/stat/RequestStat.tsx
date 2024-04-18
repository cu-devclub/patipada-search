import { RequestSummary } from "../../models/request";
import BaseStat from "./BaseStat";
import { HStack, VStack, Box, Text } from "@chakra-ui/react";
interface RequestStatProps {
  requestSummary: RequestSummary;
}
function RequestStat({ requestSummary }: RequestStatProps) {
  return (
    <Box bg='#D9D9D9' p={4} h="full">
        <BaseStat label={"Total Request"} value={requestSummary.requestAmount}>
          <HStack justify="center" spacing='4%'>
            <VStack
              w='40%'
              p='8%'
              shadow='xl' 
              borderWidth='1px' 
              bg='white'  
              alignItems="center"
              >
                <Box fontSize="xl">{requestSummary.pendingAmount}</Box>
                <Text color='gray'>Pending</Text>
            </VStack>
            <VStack
              w='40%'
              p='8%'
              shadow='md' 
              borderWidth='1px' 
              bg='white'  
              alignItems="center">
                <Box fontSize="xl">{requestSummary.reviewedAmount}</Box>
                <Text color='gray'>Reviewed</Text>
            </VStack>
          </HStack>
          </BaseStat>
    </Box>
  );
}

export default RequestStat;

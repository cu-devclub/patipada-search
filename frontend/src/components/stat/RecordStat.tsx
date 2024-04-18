import { RecordSummary } from "../../models/qa";
import BaseStat from "./BaseStat";
import { Box, HStack, Text } from "@chakra-ui/react";
interface RecordStatProps {
  recordSummary: RecordSummary;
}
function RecordStat({ recordSummary }: RecordStatProps) {
  return (
    <Box bg="gray.450" p={5} h="full">
      <BaseStat label="Total Record" value={recordSummary.recordAmount}>
        <HStack justify="center" p={2}>
          <Text>#Youtube Clip</Text>
          <Text> {recordSummary.youtubeClipAmount}</Text>
        </HStack>
      </BaseStat>
    </Box>
  );
}

export default RecordStat;

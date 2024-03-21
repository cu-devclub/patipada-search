import { RecordSummary } from "../../models/qa";
import BaseStat from "./BaseStat";
import { HStack, Text } from "@chakra-ui/react";
interface RecordStatProps {
  recordSummary: RecordSummary;
}
function RecordStat({ recordSummary }: RecordStatProps) {
  return (
    <BaseStat label={"Total Record"} value={recordSummary.recordAmount}>
      <HStack>
        <Text>#Youtube Clip</Text>
        <Text> {recordSummary.youtubeClipAmount}</Text>
      </HStack>
    </BaseStat>
  );
}

export default RecordStat;

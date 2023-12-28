import {
  Flex,
  Grid,
  VStack,
  Button,
  Box,
  Text,
  Center,
  HStack,
} from "@chakra-ui/react";
import { EditableInputForm } from "../../input";
import { Vdo } from "../../search";
import { DataItem } from "../../../models/qa";
import { useRef } from "react";
import { VdoRef } from "../../search/Vdo";
interface EditRecordFormProps {
  data: DataItem;
}
function EditRecordForm({ data }: EditRecordFormProps) {
  const vdoRef = useRef<VdoRef | null>(null);

  const handleReplay = () => {
    vdoRef.current?.replay();
  };

  return (
    <Flex flex={1} w="full" p={{ base: 2, md: 8 }} direction="column">
      <Grid
        templateColumns={{ base: "repeat(1, 1fr)", md: "60fr 40fr" }}
        gap={6}
        w="full"
        p={8}
      >
        {/* //TODO: start time, end time */}
        <VStack order={{ base: 2, md: 1 }} px={2}>
          <Grid w="full" templateColumns="auto 1fr" gap={4}>
            <Text py={1} fontWeight={"semibold"}>
              คำถาม
            </Text>
            <EditableInputForm defaultValue={data.question} />
          </Grid>
          <Grid w="full" templateColumns="auto 1fr" gap={4}>
            <Text py={1} fontWeight={"semibold"}>
              คำตอบ
            </Text>
            <EditableInputForm defaultValue={data.answer} />{" "}
          </Grid>
        </VStack>
        <VStack order={{ base: 1, md: 2 }}>
          <Box w="full">
            <Vdo
              ref={vdoRef}
              youtubeURL={data.youtubeURL}
              question={data.question}
              startTime={data.startTime}
              endTime={data.endTime}
            />
          </Box>
          <Button variant="brand_link" onClick={handleReplay}>
            เล่นอีกครั้ง
          </Button>
        </VStack>
      </Grid>
      <Center w="full">
        <HStack>
          <Button variant="success">บันทึก</Button>
          <Button variant="cancel">ยกเลิก</Button>
        </HStack>
      </Center>
    </Flex>
  );
}

export default EditRecordForm;

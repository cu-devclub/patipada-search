import {
  Flex,
  Grid,
  VStack,
  Button,
  Center,
  HStack,
  Box,
} from "@chakra-ui/react";
import { YoutubeVideo } from "../../search";
import { DataItem } from "../../../models/qa";
import { useRef } from "react";
import { VdoRef } from "../../search/YoutubeVideo";
import { TimesAndTools } from "../../search";
import React from "react";
import CommentableTextField from "./CommentableTextField";
interface EditRecordFormProps {
  data: DataItem;
}
function EditRecordForm({ data }: EditRecordFormProps) {
  const vdoRef = useRef<VdoRef | null>(null);

  const handleReplay = () => {
    vdoRef.current?.replay();
  };

  return (
    <Flex flex={1} w="full" p={{ base: 2, md: 4 }} direction="column">
      <Grid
        templateColumns={{ base: "repeat(1, 1fr)", md: "70fr 30fr" }}
        gap={6}
        w="full"
        p={2}
      >
        {/* //TODO: start time, end time */}
        <VStack order={{ base: 2, md: 1 }} px={2}>
          <CommentableTextField
            defaultValue={data.question}
            heading={"คำถาม"}
          />
          <CommentableTextField defaultValue={data.answer} heading={"คำตอบ"} />
        </VStack>
        <VStack
          order={{ base: 1, md: 2 }}
          align="center"
        >
          <Box w={{ base: "100%", lg: "60%" }}>
            <YoutubeVideo
              ref={vdoRef}
              youtubeURL={data.youtubeURL}
              question={data.question}
              startTime={data.startTime}
              endTime={data.endTime}
            />
          </Box>
          <TimesAndTools
            data={data}
            handleReplay={handleReplay}
            navigate={() => {}}
          />
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

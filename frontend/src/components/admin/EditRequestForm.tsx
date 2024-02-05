import {
  Flex,
  Grid,
  VStack,
  Button,
  Center,
  HStack,
  Box,
  useDisclosure,
} from "@chakra-ui/react";
import { YoutubeVideo } from "../search";
import { Request } from "../../models/request";
import { useRef, useState } from "react";
import { VdoRef } from "../search/YoutubeVideo";
import { TimesAndTools } from "../search";
import EditableTextField from "./EditableTextField";
import EditableTimeField from "./EditableTimeField";
import { useNavigate } from "react-router-dom";
import { checkIfCommentLeft, extractStringFromHTML } from "../../functions";
import { MessageToast } from "..";
import { ActionModal } from "../modal";
interface EditRecordFormProps {
  data: Request;
  submit: (data: Request) => void;
}
function EditRecordForm({ data, submit }: EditRecordFormProps) {
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const vdoRef = useRef<VdoRef | null>(null);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [isConfirm, setisConfirm] = useState(false);
  const modalTitle = isConfirm
    ? data.status === "reviewed"
      ? "คำขอนี้ถูกตรวจสอบแล้ว"
      : "ยืนยันคำขอแก้ไขใช่หรือไม่"
    : "ออกจากหน้าคำขอแก้ไขใช่หรือไม่";

  const modalDescription = isConfirm
    ? data.status === "reviewed"
      ? "หากยืนยันระบบจะทำการสร้างคำขอแก้ไขใหม่ ต้องการยืนยันใช่หรือไม่"
      : "หากยืนยันคำขอแก้ไขนี้จะถูกส่งให้ผู้ดูแลระบบตรวจสอบ"
    : "หากออกจากหน้าคำขอแก้ไขคำขอแก้ไขนี้จะถูกยกเลิก";

  const [startTime, setStartTime] = useState({
    isResolveComment: false,
    htmlValue: data.startTime,
    textValue: extractStringFromHTML(data.startTime),
  });
  const [endTime, setEndTime] = useState({
    isResolveComment: false,
    htmlValue: data.endTime,
    textValue: extractStringFromHTML(data.endTime),
  });
  const [question, setQuestion] = useState({
    htmlValue: data.question,
    textValue: extractStringFromHTML(data.question),
  });
  const [answer, setAnswer] = useState({
    htmlValue: data.answer,
    textValue: extractStringFromHTML(data.answer),
  });

  const handleReplay = () => {
    vdoRef.current?.replay();
  };

  const confirm = () => {
    // check if comment left
    if (
      checkIfCommentLeft(startTime.htmlValue) ||
      checkIfCommentLeft(endTime.htmlValue) ||
      checkIfCommentLeft(question.htmlValue) ||
      checkIfCommentLeft(answer.htmlValue) ||
      !startTime.isResolveComment ||
      !endTime.isResolveComment
    ) {
      addToast({
        description: "กรุณาแก้ไขความคิดเห็นทั้งหมด",
        status: "error",
      });
      onClose();
      return;
    }

    // if confirm to submit
    if (isConfirm) {
      // confirm to send request
      data.startTime = startTime.htmlValue;
      data.endTime = endTime.htmlValue;
      data.question = question.htmlValue;
      data.answer = answer.htmlValue;
      submit(data);
    } else {
      // if confirm to cancel
      // navigate back
      navigate(-1);
    }
  };

  return (
    <Flex flex={1} w="full" p={{ base: 2, md: 4 }} direction="column">
      <Grid
        templateColumns={{ base: "repeat(1, 1fr)", md: "60fr 40fr" }}
        gap={6}
        w="full"
        p={2}
      >
        <VStack order={{ base: 2, md: 1 }} px={2}>
          <EditableTimeField
            htmlValue={startTime.htmlValue}
            heading={"เวลาเริ่มต้น"}
            setHTML={(html: string) => {
              setStartTime({
                isResolveComment: startTime.isResolveComment,
                htmlValue: html,
                textValue: startTime.textValue,
              });
            }}
            setResolveComment={(isResolveComment: boolean) => {
              setStartTime({
                isResolveComment: isResolveComment,
                htmlValue: startTime.htmlValue,
                textValue: startTime.textValue,
              });
            }}
          />
          <EditableTimeField
            htmlValue={endTime.htmlValue}
            heading={"เวลาสิ้นสุด"}
            setHTML={(html: string) =>
              setEndTime({
                isResolveComment: endTime.isResolveComment,
                htmlValue: html,
                textValue: endTime.textValue,
              })
            }
            setResolveComment={(isResolveComment: boolean) => {
              setEndTime({
                isResolveComment: isResolveComment,
                htmlValue: endTime.htmlValue,
                textValue: endTime.textValue,
              });
            }}
          />
          <EditableTextField
            htmlValue={question.htmlValue}
            heading={"คำถาม"}
            setHTML={(html: string) => {
              setQuestion({
                htmlValue: html,
                textValue: question.textValue,
              });
            }}
          />
          <EditableTextField
            htmlValue={answer.htmlValue}
            heading={"คำตอบ"}
            setHTML={(html: string) =>
              setAnswer({
                htmlValue: html,
                textValue: answer.textValue,
              })
            }
          />
        </VStack>
        <VStack order={{ base: 1, md: 2 }} align="center">
          <Box w={{ base: "100%", lg: "80%" }}>
            <YoutubeVideo
              ref={vdoRef}
              youtubeURL={data.youtubeURL}
              id={data.index}
              startTime={startTime.textValue}
              endTime={endTime.textValue}
            />
          </Box>
          <TimesAndTools
            startTime={startTime.textValue}
            endTime={endTime.textValue}
            index={data.index}
            handleReplay={handleReplay}
            navigate={() => {}}
          />
        </VStack>
      </Grid>
      <Center w="full">
        <HStack>
          <Button
            variant="cancel"
            onClick={() => {
              setisConfirm(false);
              onOpen();
            }}
          >
            ยกเลิก
          </Button>
          <Button
            variant="success"
            onClick={() => {
              setisConfirm(true);
              onOpen();
            }}
          >
            บันทึก
          </Button>
        </HStack>
      </Center>
      <ActionModal
        openModal={isOpen}
        closeModal={onClose}
        modalTitle={modalTitle}
        modalBody={modalDescription}
        confirm={confirm}
      />
    </Flex>
  );
}

export default EditRecordForm;

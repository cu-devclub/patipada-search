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
import { YoutubeVideo } from "../../search";
import { Request } from "../../../models/request";
import { useRef, useState } from "react";
import { VdoRef } from "../../search/YoutubeVideo";
import { TimesAndTools } from "../../search";
import CommentableTextField from "./CommentableTextField";
import CommentableTimeField from "./CommentableTimeField";
import { useNavigate } from "react-router-dom";
import { extractStringFromHTML } from "../../../functions";
import { ActionModal } from "../../modal";
interface EditRecordFormProps {
  data: Request;
  submit: (data: Request) => void;
}

function EditRecordForm({ data, submit }: EditRecordFormProps) {
  const navigate = useNavigate();
  const vdoRef = useRef<VdoRef | null>(null);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [isConfirm, setisConfirm] = useState(false);
  const modalTitle = isConfirm
    ? "ยืนยันคำขอแก้ไขใช่หรือไม่"
    : "ออกจากหน้าคำขอแก้ไขใช่หรือไม่";
  const modalDescription = isConfirm
    ? "หากยืนยันคำขอแก้ไขนี้จะถูกส่งให้ผู้ดูแลระบบตรวจสอบ"
    : "หากออกจากหน้าคำขอแก้ไขคำขอแก้ไขนี้จะถูกยกเลิก";
  const [startTime, setStartTime] = useState({
    htmlValue: data.startTime,
  });
  const [endTime, setEndTime] = useState({
    htmlValue: data.endTime,
  });
  const [question, setQuestion] = useState({
    htmlValue: data.question,
  });
  const [answer, setAnswer] = useState({
    htmlValue: data.answer,
  });

  const handleReplay = () => {
    vdoRef.current?.replay();
  };

  const confirm = () => {
    if (isConfirm) {
      // confirm to send request
      data.startTime = startTime.htmlValue;
      data.endTime = endTime.htmlValue;
      data.question = question.htmlValue;
      data.answer = answer.htmlValue;
      onClose();
      submit(data);
    } else {
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
          <CommentableTimeField
            htmlValue={startTime.htmlValue}
            heading={"เวลาเริ่มต้น"}
            setHTML={(html: string) =>
              setStartTime({
                htmlValue: html,
              })
            }
          />
          <CommentableTimeField
            htmlValue={endTime.htmlValue}
            heading={"เวลาสิ้นสุด"}
            setHTML={(html: string) =>
              setEndTime({
                htmlValue: html,
              })
            }
          />
          <CommentableTextField
            htmlValue={question.htmlValue}
            heading={"คำถาม"}
            setHTML={(html: string) =>
              setQuestion({
                htmlValue: html,
              })
            }
          />
          <CommentableTextField
            htmlValue={answer.htmlValue}
            heading={"คำตอบ"}
            setHTML={(html: string) =>
              setAnswer({
                htmlValue: html,
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
              startTime={data.startTime}
              endTime={data.endTime}
            />
          </Box>
          <TimesAndTools
            index={data.index}
            startTime={extractStringFromHTML(data.startTime)}
            endTime={extractStringFromHTML(data.endTime)}
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
        confirm={confirm}
        modalTitle={modalTitle}
        modalBody={modalDescription}
      />
    </Flex>
  );
}

export default EditRecordForm;

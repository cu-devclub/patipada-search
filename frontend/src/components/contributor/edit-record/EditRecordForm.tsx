import {
  Flex,
  Grid,
  VStack,
  Button,
  Center,
  HStack,
  Box,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
} from "@chakra-ui/react";
import { YoutubeVideo } from "../../search";
import { Request } from "../../../models/request";
import { useRef, useState } from "react";
import { VdoRef } from "../../search/YoutubeVideo";
import { TimesAndTools } from "../../search";
import React from "react";
import CommentableTextField from "./CommentableTextField";
import CommentableTimeField from "./CommentableTimeField";
import { extractStringFromHTML } from "../../../functions";
import { useNavigate } from "react-router-dom";
interface EditRecordFormProps {
  data: Request;
  submit: (data: Request) => void;
}
function EditRecordForm({ data, submit }: EditRecordFormProps) {
  const navigate = useNavigate();
  const vdoRef = useRef<VdoRef | null>(null);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [isConfirm, setisConfirm] = useState(false);
  const [isSomeFormEditing, setSomeFormEditing] = useState(false)
  const modalTitle = isConfirm
    ? "ยืนยันคำขอแก้ไขใช่หรือไม่"
    : "ออกจากหน้าคำขอแก้ไขใช่หรือไม่";
  const modalDescription = isConfirm
    ? "หากยืนยันคำขอแก้ไขนี้จะถูกส่งให้ผู้ดูแลระบบตรวจสอบ"
    : "หากออกจากหน้าคำขอแก้ไขคำขอแก้ไขนี้จะถูกยกเลิก";
  const [startTime, setStartTime] = useState({
    htmlValue: data.startTime,
    textValue: extractStringFromHTML(data.startTime),
  });
  const [endTime, setEndTime] = useState({
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
    if (isConfirm) {
      // confirm to send request
      data.startTime = startTime.htmlValue;
      data.endTime = endTime.htmlValue;
      data.question = question.htmlValue;
      data.answer = answer.htmlValue;
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
            textValue={startTime.textValue}
            heading={"เวลาเริ่มต้น"}
            isOtherFormEditing={isSomeFormEditing}
            setFormEditing={setSomeFormEditing}
            confirm={(html: string) =>
              setStartTime({
                htmlValue: html,
                textValue: extractStringFromHTML(html),
              })
            }
          />
          <CommentableTimeField
            htmlValue={endTime.htmlValue}
            textValue={endTime.textValue}
            heading={"เวลาสิ้นสุด"}
            isOtherFormEditing={isSomeFormEditing}
            setFormEditing={setSomeFormEditing}
            confirm={(html: string) => {
              setEndTime({
                htmlValue: html,
                textValue: extractStringFromHTML(html),
              });
            }}
          />
          <CommentableTextField
            htmlValue={question.htmlValue}
            textValue={question.textValue}
            heading={"คำถาม"}
            isOtherFormEditing={isSomeFormEditing}
            setFormEditing={setSomeFormEditing}
            confirm={(html: string) =>
              setQuestion({
                htmlValue: html,
                textValue: extractStringFromHTML(html),
              })
            }
          />
          <CommentableTextField
            htmlValue={answer.htmlValue}
            textValue={answer.textValue}
            heading={"คำตอบ"}
            isOtherFormEditing={isSomeFormEditing}
            setFormEditing={setSomeFormEditing}
            confirm={(html: string) =>
              setAnswer({
                htmlValue: html,
                textValue: extractStringFromHTML(html),
              })
            }
          />
        </VStack>
        <VStack order={{ base: 1, md: 2 }} align="center">
          <Box w={{ base: "100%", lg: "80%" }}>
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
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>{modalTitle}</ModalHeader>
          <ModalCloseButton />
          <ModalBody>{modalDescription}</ModalBody>

          <ModalFooter>
            <Button variant="cancel" mr={3} onClick={onClose}>
              ยกเลิก
            </Button>
            <Button variant="success" onClick={confirm}>
              ยืนยัน
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Flex>
  );
}

export default EditRecordForm;

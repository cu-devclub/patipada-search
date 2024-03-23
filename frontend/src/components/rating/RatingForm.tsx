import {
  HStack,
  Text,
  useRadioGroup,
  Center,
  Heading,
  VStack,
  Textarea,
  Button,
  Link,
} from "@chakra-ui/react";
import { useState } from "react";
import RadioCard from "./RadioCard";
import { MessageToast } from "..";
import { Rating } from "../../models/ratings";
import { insertRatingsService } from "../../service/data";
import { ExternalLinkIcon } from "@chakra-ui/icons";

interface RatingFormProps {
  closeModal: () => void;
}

function RatingForm({ closeModal }: RatingFormProps) {
  const { addToast } = MessageToast();

  const options = ["1", "2", "3", "4", "5"];

  const [rating, setRating] = useState("1");

  const [comment, setComment] = useState("");

  const { getRootProps, getRadioProps } = useRadioGroup({
    name: "radio",
    defaultValue: "1",
    onChange: setRating,
    value: rating,
  });

  const group = getRootProps();

  const submit = async () => {
    const ratingModel: Rating = {
      stars: parseInt(rating),
      comment: comment,
    };

    await insertRatingsService(ratingModel)
      .then(() => {
        addToast({
          description: "ข้อเสนอแนะของคุณถูกบันทึกเรียบร้อยแล้ว",
          status: "success",
        });
        closeModal();
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดในการบันทึกข้อมูล",
          status: "error",
        });
      });
  };

  return (
    <VStack w="full" borderRadius="2rem">
      <Heading color="black">ข้อเสนอแนะ</Heading>
      <Text fontSize="md" color="gray.400" as="p">
        ให้คะแนนระบบของเรา และความพึงพอใจจากการใช้งานเว็บไซต์
      </Text>

      <Center color="white">
        <HStack {...group}>
          {options.map((value) => {
            const radio = getRadioProps({ value });
            return (
              <RadioCard key={value} {...radio}>
                {value}
              </RadioCard>
            );
          })}
        </HStack>
      </Center>

      <Textarea
        mt={4}
        placeholder={"ความคิดเห็น"}
        w="100%"
        value={comment}
        onChange={(e) => setComment(e.target.value)}
      />

      <Text>
        <Text as="u">หรือ</Text> ทำแบบประเมินโดยละเอียดผ่าน{" "}
        <Link href="https://chakra-ui.com" isExternal color="blue.600">
          Google form <ExternalLinkIcon mx="2px" />
        </Link>
      </Text>

      <Button colorScheme="orange" mt={2} onClick={submit}>
        ส่งข้อเสนอแนะ
      </Button>
    </VStack>
  );
}

export default RatingForm;

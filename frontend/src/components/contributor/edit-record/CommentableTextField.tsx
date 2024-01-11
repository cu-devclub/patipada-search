import { VStack, Text } from "@chakra-ui/react";
import React from "react";
import { CommentTiptap } from "../../Tiptap";
interface CommentableTextFieldProps {
  heading: string;
  htmlValue: string;
  setHTML: (html: string) => void;
}

function CommentableTextField({
  heading,
  htmlValue,
  setHTML,
}: CommentableTextFieldProps) {
  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>

      <CommentTiptap defaultValue={htmlValue} setHTML={setHTML} />
    </VStack>
  );
}

export default CommentableTextField;

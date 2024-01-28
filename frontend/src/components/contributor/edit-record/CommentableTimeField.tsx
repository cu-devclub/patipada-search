import { VStack, Text } from "@chakra-ui/react";
import { TimeCommentTiptap } from "../../Tiptap";
interface CommentableTimeFieldProps {
  heading: string;
  htmlValue: string;
  setHTML: (html: string) => void;
}

function CommentableTimeField({
  heading,
  htmlValue,
  setHTML,
}: CommentableTimeFieldProps) {
  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>
      <TimeCommentTiptap defaultValue={htmlValue} setHTML={setHTML} />
    </VStack>
  );
}

export default CommentableTimeField;

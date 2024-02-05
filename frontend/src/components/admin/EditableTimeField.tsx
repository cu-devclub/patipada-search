import { VStack, Text } from "@chakra-ui/react";
import { EditableTimeTiptap } from "../Tiptap";
interface EditableTimeFieldProps {
  heading: string;
  htmlValue: string;
  setHTML: (html: string) => void;
  setResolveComment: (isResolveComment: boolean) => void;
}

function EditableTimeField({
  heading,
  htmlValue,
  setHTML,
  setResolveComment,
}: EditableTimeFieldProps) {
  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>
      <EditableTimeTiptap
        defaultValue={htmlValue}
        setHTML={setHTML}
        setResolveComment={setResolveComment}
      />
    </VStack>
  );
}

export default EditableTimeField;
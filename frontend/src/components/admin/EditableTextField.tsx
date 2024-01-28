import { VStack, Text } from "@chakra-ui/react";
import { EditableTiptap } from "../Tiptap";
interface EditableTextFieldProps {
  heading: string;
  htmlValue: string;
  setHTML: (html: string) => void;
}

function EditableTextField({
  heading,
  htmlValue,
  setHTML,
}: EditableTextFieldProps) {
  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>

      <EditableTiptap defaultValue={htmlValue} setHTML={setHTML} />
    </VStack>
  );
}

export default EditableTextField;

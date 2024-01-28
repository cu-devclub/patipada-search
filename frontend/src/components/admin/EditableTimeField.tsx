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

// 🚀 ~ htmlValue: <p><span data-comment="{&quot;uuid&quot;:&quot;3a88707f-4189-42c7-87d6-2c53a2c6376f&quot;,&quot;comments&quot;:[{&quot;userName&quot;:&quot;super-admin&quot;,&quot;time&quot;:1705376318570,&quot;content&quot;:&quot;00:31:17&quot;}]}">00:29:17</span></p>

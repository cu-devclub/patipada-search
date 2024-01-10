import { VStack, Tooltip, Text } from "@chakra-ui/react";
import React, { useState } from "react";
import { TimeCommentTiptap } from "../../Tiptap";
interface CommentableTimeFieldProps {
  heading: string;
  textValue: string;
  htmlValue: string;
  isOtherFormEditing;
  setFormEditing;
  confirm: (html: string) => void;
}

function CommentableTimeField({
  heading,
  textValue,
  htmlValue,
  isOtherFormEditing,
  setFormEditing,
  confirm,
}: CommentableTimeFieldProps) {
  const [edit, setedit] = useState(false);
  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>
      {edit ? (
        <TimeCommentTiptap
          defaultValue={htmlValue}
          cancel={() => {
            setedit(false);
            setFormEditing(false);
          }}
          confirm={confirm}
        />
      ) : (
        <Tooltip label="click to highlight">
          <Text
            py={1}
            fontWeight={"light"}
            _hover={{
              background: "gray.200",
            }}
            onClick={() => {
              if (isOtherFormEditing) return;
              setedit(true);
              setFormEditing(true);
            }}
          >
            {textValue}
          </Text>
        </Tooltip>
      )}
    </VStack>
  );
}

export default CommentableTimeField;

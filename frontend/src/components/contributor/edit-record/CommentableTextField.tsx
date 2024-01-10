import { VStack, Tooltip, Text } from "@chakra-ui/react";
import React, { useState } from "react";
import { CommentTiptap } from "../../Tiptap";
interface CommentableTextFieldProps {
  heading: string;
  textValue: string;
  htmlValue: string;
  isOtherFormEditing: boolean;
  setFormEditing: (isEditing: boolean) => void;
  confirm: (html: string) => void;
}

function CommentableTextField({
  heading,
  textValue,
  htmlValue,
  isOtherFormEditing,
  setFormEditing,
  confirm,
}: CommentableTextFieldProps) {
  const [edit, setedit] = useState(false);
  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>
      {edit  ? (
        <CommentTiptap
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

export default CommentableTextField;

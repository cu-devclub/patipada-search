import { VStack, Tooltip, Text } from "@chakra-ui/react";
import React, { useState } from "react";
import { CommentTiptap } from "../../Tiptap";

interface CommentableTextFieldProps {
  heading: string;
  defaultValue: string;
}

function CommentableTextField({
  heading,
  defaultValue,
}: CommentableTextFieldProps) {
  const [edit, setedit] = useState(false);

  return (
    <VStack w="full" spacing={0} align="start">
      <Text py={1} fontWeight={"semibold"}>
        {heading}
      </Text>
      {edit ? (
        <CommentTiptap
          defaultValue={defaultValue}
          cancel={() => setedit(false)}
        />
      ) : (
        <Tooltip label="click to highlight">
          <Text
            py={1}
            fontWeight={"light"}
            _hover={{
              background: "gray.200",
            }}
            onClick={() => setedit(true)}
          >
            {defaultValue}
          </Text>
        </Tooltip>
      )}
    </VStack>
  );
}

export default CommentableTextField;

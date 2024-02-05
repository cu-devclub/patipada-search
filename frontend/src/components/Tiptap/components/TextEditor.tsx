import { Box } from "@chakra-ui/react";
import { ReactNode } from "react";
import "../Tiptap.scss";

interface TextEditorProps {
  children: ReactNode;
}

function TextEditor({ children }: TextEditorProps) {
  return (
    <Box className="tiptap" w="100%">
      <Box className="tiptap-Box" fontWeight={"normal"}>
        {children}
      </Box>
    </Box>
  );
}

export default TextEditor;

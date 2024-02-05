/* eslint-disable @typescript-eslint/no-explicit-any */
import { HStack, Button } from "@chakra-ui/react";
import { BubbleMenu, Editor,EditorContent } from "@tiptap/react";
import { ReactNode } from "react";

interface CommentBubbleProps {
  editor: Editor | null;
  setCommentText: (text: string) => void;
  setComment: () => void;
  inputField: ReactNode;
}

function CommentBubble({
  editor,
  setCommentText,
  setComment,
  inputField,
}: CommentBubbleProps) {
  return (
    <>
      {editor && (
        <BubbleMenu
          tippy-options={{ duration: 100, placement: "bottom" }}
          editor={editor}
          shouldShow={() => !editor?.view.state.selection.empty}
        >
          {inputField}
          <HStack>
            <Button onClick={() => setCommentText("")} variant="cancel">
              Clear
            </Button>
            <Button onClick={() => setComment()} variant="success">
              Add
            </Button>
          </HStack>
        </BubbleMenu>
      )}
      <EditorContent className="editor-content" editor={editor} />
    </>
  );
}

export default CommentBubble;

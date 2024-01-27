/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import * as React from "react";
import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import { format } from "date-fns/format";
import "../Tiptap.scss";
import { Comment } from "../extensions/comment";
import {  Box, Flex, HStack, Text } from "@chakra-ui/react";
const dateTimeFormat = "dd.MM.yyyy HH:mm";

interface CommentInstance {
  uuid?: string;
  comments?: any[];
}

interface TipTapProps {
  defaultValue: string;
  setHTML: (html: string) => void;
}

const EditableTiptap = ({ defaultValue, setHTML }: TipTapProps) => {
  const editor = useEditor({
    extensions: [StarterKit, Comment],
    content: defaultValue || "",
    editable: true,
    onUpdate({ editor }) {
      findCommentsAndStoreValues();

      setCurrentComment(editor);
      setHTML(editor?.getHTML() || defaultValue);
    },

    onSelectionUpdate({ editor }) {
      setCurrentComment(editor);

      setIsTextSelected(!!editor.state.selection.content().size);
    },

    editorProps: {
      attributes: {
        spellcheck: "false",
      },
    },
  });

  const [, setShowCommentMenu] = React.useState(false);

  const [, setIsTextSelected] = React.useState(false);

  const [, setShowAddCommentSection] = React.useState(true);

  const formatDate = (d: any) =>
    d ? format(new Date(d), dateTimeFormat) : null;

  const [activeCommentsInstance, setActiveCommentsInstance] =
    React.useState<CommentInstance>({});

  const [allComments, setAllComments] = React.useState<any[]>([]);

  const findCommentsAndStoreValues = () => {
    const parser = new DOMParser();
    const htmlText = editor?.getHTML() || defaultValue;
    const doc = parser.parseFromString(htmlText, "text/html");
    const comments = doc.querySelectorAll("span[data-comment]");

    const tempComments: any[] = [];

    comments.forEach((node) => {
      const nodeComments = node.getAttribute("data-comment");
      const jsonComments = nodeComments ? JSON.parse(nodeComments) : null;

      if (jsonComments !== null) {
        tempComments.push({
          node,
          jsonComments,
        });
      }
    });

    setAllComments(tempComments);
  };
  const setCurrentComment = (editor: any) => {
    const newVal = editor.isActive("comment");

    if (newVal) {
      setTimeout(() => setShowCommentMenu(newVal), 50);

      setShowAddCommentSection(!editor.state.selection.empty);

      const parsedComment = JSON.parse(editor.getAttributes("comment").comment);

      parsedComment.comment =
        typeof parsedComment.comments === "string"
          ? JSON.parse(parsedComment.comments)
          : parsedComment.comments;

      setActiveCommentsInstance(parsedComment);
    } else {
      setActiveCommentsInstance({});
    }
  };

  React.useEffect(() => {
    const timeoutId = setTimeout(findCommentsAndStoreValues, 100);
    return () => clearTimeout(timeoutId); // This is the cleanup function
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <Flex
      dir="row"
      className="tiptap"
      w="100%"
      bg="gray.100"
      p={2}
      borderRadius={"lg"}
      position="relative"
    >
      <Box className="tiptap-Box" w="60%" fontWeight={"normal"}>
        <EditorContent className="editor-content" editor={editor} />
      </Box>

      <Flex direction="column" pb={10}>
        {allComments.map((comment, i) => {
          return (
            <Box
              key={i + "external_comment"}
              bg="gray.100"
              shadow="lg"
              my={2}
              borderRadius={"md"}
              w="sm"
            >
              {comment.jsonComments.comments.map(
                (jsonComment: any, j: number) => {
                  return (
                    <Box
                      key={`${j}_${Math.random()}`}
                      p={3}
                      border={
                        comment.jsonComments.uuid ===
                        activeCommentsInstance.uuid
                          ? "2px"
                          : "none"
                      }
                      borderColor={
                        comment.jsonComments.uuid ===
                        activeCommentsInstance.uuid
                          ? "red.500"
                          : "gray.300"
                      }
                    >
                      <Flex direction="column">
                        <HStack>
                          <Text fontWeight={"semibold"}>
                            {jsonComment.userName}
                          </Text>
                          <Text fontSize={"sm"}>
                            {formatDate(jsonComment.time)}
                          </Text>
                        </HStack>
                        <Text>{jsonComment.content}</Text>
                      </Flex>
                    </Box>
                  );
                }
              )}
            </Box>
          );
        })}
      </Flex>
    </Flex>
  );
};

export default EditableTiptap;

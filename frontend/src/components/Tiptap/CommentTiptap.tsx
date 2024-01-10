/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import * as React from "react";
import { useEditor, EditorContent, BubbleMenu } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import { format } from "date-fns/format";
import "./Tiptap.scss";
import { Comment } from "./extensions/comment";
import { v4 as uuidv4 } from "uuid";
import {
  Button,
  Box,
  Flex,
  Textarea,
  HStack,
  Text,
  ButtonGroup,
} from "@chakra-ui/react";
// import { setTimeout } from "../../functions/time";
import { getCookie } from "typescript-cookie";
const dateTimeFormat = "dd.MM.yyyy HH:mm";

interface CommentInstance {
  uuid?: string;
  comments?: any[];
}

interface TipTapProps {
  defaultValue: string;
  cancel: () => void;
  confirm: (html: string) => void;
}

const CommentTiptap = ({ defaultValue, cancel, confirm }: TipTapProps) => {
  const username = getCookie("username");
  const editor = useEditor({
    extensions: [StarterKit, Comment],
    content: defaultValue || "",
    onUpdate({ editor }) {
      findCommentsAndStoreValues();

      setCurrentComment(editor);
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

  const [commentText, setCommentText] = React.useState("");

  const [showCommentMenu, setShowCommentMenu] = React.useState(false);

  const [isTextSelected, setIsTextSelected] = React.useState(false);

  const [showAddCommentSection, setShowAddCommentSection] =
    React.useState(true);

  const formatDate = (d: any) =>
    d ? format(new Date(d), dateTimeFormat) : null;

  const [activeCommentsInstance, setActiveCommentsInstance] =
    React.useState<CommentInstance>({});

  const [allComments, setAllComments] = React.useState<any[]>([]);

  const findCommentsAndStoreValues = () => {
    const proseMirror = document.querySelector(".ProseMirror");

    const comments = proseMirror?.querySelectorAll("span[data-comment]");

    const tempComments: any[] = [];

    if (!comments) {
      setAllComments([]);
      return;
    }

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

  const setComment = () => {
    if (!commentText.trim().length) return;

    const activeCommentInstance: CommentInstance = JSON.parse(
      JSON.stringify(activeCommentsInstance)
    );

    const commentsArray =
      typeof activeCommentInstance.comments === "string"
        ? JSON.parse(activeCommentInstance.comments)
        : activeCommentInstance.comments;

    if (commentsArray) {
      commentsArray.push({
        userName: username,
        time: Date.now(),
        content: commentText,
      });

      const commentWithUuid = JSON.stringify({
        uuid: activeCommentsInstance.uuid || uuidv4(),
        comments: commentsArray,
      });

      // eslint-disable-next-line no-unused-expressions
      editor?.chain().setComment(commentWithUuid).run();
    } else {
      const commentWithUuid = JSON.stringify({
        uuid: uuidv4(),
        comments: [
          {
            userName: username,
            time: Date.now(),
            content: commentText,
          },
        ],
      });

      // eslint-disable-next-line no-unused-expressions
      editor?.chain().setComment(commentWithUuid).run();
    }

    setTimeout(() => setCommentText(""), 0.1);

    // force user to unselect
    editor?.commands.focus(editor?.state.doc.content.size);
  };

  const submit = () => {
    confirm(editor?.getHTML() || defaultValue);
    cancel();
  };

  React.useEffect(() => {
    const timeoutId = setTimeout(findCommentsAndStoreValues, 100);
    return () => clearTimeout(timeoutId); // This is the cleanup function
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
      <Box className="tiptap-Box" w="60%">
        {editor && (
          <BubbleMenu
            tippy-options={{ duration: 100, placement: "bottom" }}
            editor={editor}
            shouldShow={() => !editor?.view.state.selection.empty}
          >
            <Textarea
              value={commentText}
              onChange={(e) => setCommentText((e.target as any).value)}
              placeholder="Add comment..."
              bg="white"
              mb={2}
            />
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
                      borderBottom="2px"
                      borderColor="gray.300"
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

              {comment.jsonComments.uuid === activeCommentsInstance.uuid && (
                <Flex w="full" direction="column" gap={1}>
                  <Textarea
                    value={commentText}
                    onChange={(e) => setCommentText((e.target as any).value)}
                    placeholder="Add comment..."
                    bg="white"
                  />

                  <HStack>
                    <Button onClick={() => setCommentText("")} variant="cancel">
                      Clear
                    </Button>
                    <Button onClick={() => setComment()} variant="success">
                      Add
                    </Button>
                  </HStack>
                </Flex>
              )}
            </Box>
          );
        })}
      </Flex>
      <ButtonGroup position="absolute" bottom={0} right={0}>
        <Button variant="cancel" onClick={cancel}>
          ยกเลิก
        </Button>
        <Button variant="success" onClick={submit}>
          ยืนยัน
        </Button>
      </ButtonGroup>
    </Flex>
  );
};

export default CommentTiptap;

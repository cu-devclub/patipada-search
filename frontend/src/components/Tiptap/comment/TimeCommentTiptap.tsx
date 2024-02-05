/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState, useEffect } from "react";
import { EditorContent } from "@tiptap/react";
import "../Tiptap.scss";
import { Box, Flex } from "@chakra-ui/react";
import { getCookie } from "typescript-cookie";
import { extractStringFromHTML, splitTime } from "../../../functions";
import {
  CommentBubble,
  CommentCardWithDiscussion,
  TimeInput,
} from "../components";
import {
  findCommentsAndStoreValue,
  handleChangeTimeStamp,
  setComment,
} from "../../../functions/tiptap";
import { CommentInstance } from "../../../models/comment";
import { GetEditor } from "./editor";
import { TimeStamps } from "../../../models/time";

interface TipTapProps {
  defaultValue: string;
  setHTML: (html: string) => void;
}
const TimeCommentTiptap = ({ defaultValue, setHTML }: TipTapProps) => {
  const { hours, minutes, seconds } = splitTime(
    extractStringFromHTML(defaultValue)
  );

  const [timeStamp, setTimeStamp] = useState<TimeStamps>({
    hours: hours,
    minutes: minutes,
    seconds: seconds,
  });

  const username = getCookie("username");

  const [commentText, setCommentText] = useState("");

  const [activeCommentsInstance, setActiveCommentsInstance] =
    useState<CommentInstance>({});

  const [allComments, setAllComments] = useState<any[]>([]);

  const setCommentInherit = () => {
    setComment(
      editor,
      activeCommentsInstance,
      username || "",
      commentText,
      setCommentText,
      setHTML
    );
  };

  const editor = GetEditor(
    defaultValue,
    setAllComments,
    setActiveCommentsInstance
  );

  const updateNewTimeStamp = (e: number, action: string) => {
    handleChangeTimeStamp(e, action, setTimeStamp, setCommentText);
  };

  const timeState = (
    <TimeInput timeStamp={timeStamp} setTimeStamp={updateNewTimeStamp} />
  );

  const inputField = <Flex p="2">{timeState}</Flex>;

  const findComment = () => {
    findCommentsAndStoreValue(
      editor?.getHTML() || defaultValue,
      setAllComments
    );
  };

  useEffect(() => {
    const timeoutId = setTimeout(findComment, 100);
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
      <Box className="tiptap-Box" w="60%">
        <CommentBubble
          editor={editor}
          setCommentText={setCommentText}
          setComment={setCommentInherit}
          inputField={inputField}
        />
        <EditorContent className="editor-content" editor={editor} />
      </Box>
      <CommentCardWithDiscussion
        allComments={allComments}
        commentText={commentText}
        setCommentText={setCommentText}
        activeCommentInstance={activeCommentsInstance}
        setComment={setCommentInherit}
        inputField={inputField}
      />
    </Flex>
  );
};

export default TimeCommentTiptap;

/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState, useEffect } from "react";
import { useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import "../Tiptap.scss";
import { Comment } from "../extensions/comment";
import {
  splitTime,
  extractStringFromHTML,
} from "../../../functions";
import { CommentCard, Layout, TextEditor, TimeInput } from "../components";
import {
  findCommentsAndStoreValue,
  handleChangeTimeStamp,
} from "../../../functions/tiptap";
import { TimeStamps } from "../../../models/time";

interface TipTapProps {
  defaultValue: string;
  setHTML: (html: string) => void;
  setResolveComment: (isResolveComment: boolean) => void;
}
const EditableTimeTiptap = ({
  defaultValue,
  setHTML,
  setResolveComment,
}: TipTapProps) => {
  const { hours, minutes, seconds } = splitTime(
    extractStringFromHTML(defaultValue)
  );

  const [timeStamp, setTimeStamp] = useState<TimeStamps>({
    hours: hours,
    minutes: minutes,
    seconds: seconds,
  });

  const updateNewTimeStamp = (e: number, action: string) => {
    handleChangeTimeStamp(e, action, setTimeStamp, setHTML);
  };

  const editor = useEditor({
    extensions: [StarterKit, Comment],
    content: defaultValue || "",
    editable: true,

    editorProps: {
      attributes: {
        spellcheck: "false",
      },
    },
  });

  const [allComments, setAllComments] = useState<any[]>([]);

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

  useEffect(() => {
    setResolveComment(allComments.length == 0);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [allComments]);

  const resolveComment = (comment: any) => {
    const newComments = allComments.filter(
      (c) => c.jsonComments !== comment.jsonComments
    );
    setAllComments(newComments);
  };

  return (
    <Layout>
      <TextEditor>
        <TimeInput timeStamp={timeStamp} setTimeStamp={updateNewTimeStamp} />
      </TextEditor>
      <CommentCard allComments={allComments} resolveComment={resolveComment} />
    </Layout>
  );
};

export default EditableTimeTiptap;

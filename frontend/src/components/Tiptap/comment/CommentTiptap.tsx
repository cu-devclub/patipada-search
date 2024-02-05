/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import { EditorContent } from "@tiptap/react";
import "../Tiptap.scss";
import { Textarea } from "@chakra-ui/react";
import { getCookie } from "typescript-cookie";
import {
  CommentBubble,
  CommentCardWithDiscussion,
  TextEditor,
  Layout,
} from "../components";
import {
  findCommentsAndStoreValue,
  setComment,
} from "../../../functions/tiptap";
import { GetEditor } from "./editor";
import { CommentInstance } from "../../../models/comment";

interface TipTapProps {
  defaultValue: string;
  setHTML: (html: string) => void;
}

const CommentTiptap = ({ defaultValue, setHTML }: TipTapProps) => {
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

  const inputField = (
    <Textarea
      value={commentText}
      onChange={(e) => setCommentText((e.target as any).value)}
      placeholder="Add comment..."
      bg="white"
    />
  );

  const editor = GetEditor(
    defaultValue,
    setAllComments,
    setActiveCommentsInstance
  );

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
    <Layout>
      {editor && (
        <TextEditor>
          <CommentBubble
            editor={editor}
            setCommentText={setCommentText}
            setComment={setCommentInherit}
            inputField={inputField}
          />
          <EditorContent className="editor-content" editor={editor} />
        </TextEditor>
      )}
      <CommentCardWithDiscussion
        allComments={allComments}
        commentText={commentText}
        setCommentText={setCommentText}
        activeCommentInstance={activeCommentsInstance}
        setComment={setCommentInherit}
        inputField={inputField}
      />
    </Layout>
  );
};

export default CommentTiptap;

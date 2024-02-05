/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState, useEffect } from "react";
import { EditorContent, useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import "../Tiptap.scss";
import { Comment } from "../extensions/comment";
import {
  getStartAndEndIndexOfComments,
  removeCommentFromHTML,
} from "../../../functions";
import { CommentCard, TextEditor, Layout } from "../components";
import { setCurrentComment } from "../../../functions/tiptap";

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
      setCurrentComment(editor, setActiveCommentsInstance);
      setHTML(editor?.getHTML() || defaultValue);
    },
    onSelectionUpdate({ editor }) {
      setCurrentComment(editor, setActiveCommentsInstance);
    },

    editorProps: {
      attributes: {
        spellcheck: "false",
      },
    },
  });

  const [activeCommentsInstance, setActiveCommentsInstance] =
    useState<CommentInstance>({});

  const [allComments, setAllComments] = useState<any[]>([]);

  const findCommentsAndStoreValues = () => {
    const parser = new DOMParser();
    const htmlText = editor?.getHTML() || defaultValue;
    const doc = parser.parseFromString(htmlText, "text/html");
    const comments = doc.querySelectorAll("span[data-comment]");
    const tempComments: any[] = [];
    const commentPosition = getStartAndEndIndexOfComments(htmlText);
    let indexCounter = 0;
    comments.forEach((node) => {
      const nodeComments = node.getAttribute("data-comment");
      const jsonComments = nodeComments ? JSON.parse(nodeComments) : null;

      if (jsonComments !== null) {
        const start = commentPosition[indexCounter][0];
        const end = commentPosition[indexCounter][1];
        tempComments.push({
          node,
          jsonComments,
          start,
          end,
        });
        indexCounter++;
      }
    });

    setAllComments(tempComments);
  };

  useEffect(() => {
    const timeoutId = setTimeout(findCommentsAndStoreValues, 100);
    return () => clearTimeout(timeoutId); // This is the cleanup function
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (activeCommentsInstance.uuid) {
    // user focus on some comment
    // check if cursor at the end of the comment
    // if yes, then unset comment mark
    const cursorPosition = editor?.state.selection.from;
    for (let i = 0; i < allComments.length; i++) {
      if (cursorPosition === allComments[i].end) {
        editor?.commands.unsetMark("comment");
        break;
      }
    }
  }

  const resolveComment = (comment: any, time: string) => {
    const newHTML = removeCommentFromHTML(
      editor?.getHTML() || defaultValue,
      comment.start.index,
      comment.end.index,
      time
    );

    editor?.commands.setContent(newHTML);
    setHTML(newHTML);

    const newComments = allComments.flatMap((c) => {
      if (c.jsonComments.comments.length > 1) {
        const updatedComments = c.jsonComments.comments.filter(
          (com: any) => com.time !== time
        );
        return { ...c, jsonComments: { ...c.jsonComments, comments: updatedComments } };
      } else if (c.jsonComments.uuid === comment.jsonComments.uuid) {
        return [];
      } else {
        return c;
      }
    });

    setAllComments(newComments);
    findCommentsAndStoreValues();
  };

  return (
    <Layout>
      <TextEditor>
        <EditorContent className="editor-content" editor={editor} />
      </TextEditor>
      <CommentCard allComments={allComments} resolveComment={resolveComment} />
    </Layout>
  );
};

export default EditableTiptap;


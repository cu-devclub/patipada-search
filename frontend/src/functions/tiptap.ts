/* eslint-disable @typescript-eslint/no-explicit-any */
import { v4 as uuidv4 } from "uuid";
import { CommentInstance } from "../models/comment";
import { generateTime } from "./time";
import { TimeStamps } from "../models/time";

export const findCommentsAndStoreValue = (
  htmlText: string,
  setAllComments: any
) => {
  const parser = new DOMParser();
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

export const setCurrentComment = (editor: any, setActiveComment: any) => {
  const newVal = editor.isActive("comment");

  if (newVal) {
    const parsedComment = JSON.parse(editor.getAttributes("comment").comment);

    parsedComment.comment =
      typeof parsedComment.comments === "string"
        ? JSON.parse(parsedComment.comments)
        : parsedComment.comments;

    setActiveComment(parsedComment);
  } else {
    setActiveComment({});
  }
};

export const setComment = (
  editor: any,
  comment: any,
  username: string,
  commentText: string,
  setCommentText: any,
  setHTML: any
) => {
  if (!commentText.trim().length) return;

  const activeCommentInstance: CommentInstance = JSON.parse(
    JSON.stringify(comment)
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
      uuid: activeCommentInstance.uuid || uuidv4(),
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

  setHTML(editor?.getHTML());
};

export const handleChangeTimeStamp = (
    e: number,
    action: string,
    setTimeStamp: any,
    setRelatedText: any
) => {
    setTimeStamp((prevTimeStamp: TimeStamps) => {
        const updatedTimeStamp = { ...prevTimeStamp, [action]: e };
        const fullTimeText = generateTime(
            updatedTimeStamp.hours,
            updatedTimeStamp.minutes,
            updatedTimeStamp.seconds
        );

        setRelatedText(fullTimeText);

        return updatedTimeStamp;
    });
};
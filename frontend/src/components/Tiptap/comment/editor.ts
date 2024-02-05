/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import { Comment } from "../extensions/comment";
import {
  findCommentsAndStoreValue,
  setCurrentComment,
} from "../../../functions/tiptap";

export const GetEditor = (defaultValue:string,setAllComments:any,setActiveCommentsInstance:any) => {
    const editor = useEditor({
      extensions: [StarterKit, Comment],
      content: defaultValue || "",
      editable: false,
      onUpdate({ editor }) {
        findCommentsAndStoreValue(
          editor?.getHTML() || defaultValue,
          setAllComments
        );

        setCurrentComment(editor, setActiveCommentsInstance);
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

    return editor;
}
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import * as React from "react";
import { useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import { format } from "date-fns/format";
import "../Tiptap.scss";
import { Comment } from "../extensions/comment";
import {
  Button,
  Box,
  Flex,
  HStack,
  Text,
  NumberInput,
  NumberDecrementStepper,
  NumberIncrementStepper,
  NumberInputField,
  NumberInputStepper,
} from "@chakra-ui/react";
import {
  splitTime,
  generateTime,
  extractStringFromHTML,
} from "../../../functions";
const dateTimeFormat = "dd.MM.yyyy HH:mm";


interface TipTapProps {
  defaultValue: string;
  setHTML: (html: string) => void;
  setResolveComment: (isResolveComment: boolean) => void;
}
const EditableTimeTiptap = ({ defaultValue, setHTML,setResolveComment }: TipTapProps) => {
  const { hours, minutes, seconds } = splitTime(
    extractStringFromHTML(defaultValue)
  );
  const [hourState, setHourState] = React.useState(hours);
  const [minuteState, setMinuteState] = React.useState(minutes);
  const [secondState, setSecondState] = React.useState(seconds);
  const handleChangeHour = (e: any) => {
    setHourState(e);
    const fullTimeText = generateTime(e, minuteState, secondState);
    setHTML(fullTimeText);
  };

  const handleChangeMinute = (e: any) => {
    setMinuteState(e);
    const fullTimeText = generateTime(hourState, e, secondState);
    setHTML(fullTimeText);
  };

  const handleChangeSecond = (e: any) => {
    setSecondState(e);
    const fullTimeText = generateTime(hourState, minuteState, e);
    setHTML(fullTimeText);
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

  const formatDate = (d: any) =>
    d ? format(new Date(d), dateTimeFormat) : null;

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

  React.useEffect(() => {
    const timeoutId = setTimeout(findCommentsAndStoreValues, 100);
    return () => clearTimeout(timeoutId); // This is the cleanup function
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  React.useEffect(()=>{
    setResolveComment(allComments.length == 0);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  },[allComments])

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
        <HStack shouldWrapChildren mb={2}>
          <NumberInput
            bg="white"
            size="xs"
            maxW={16}
            min={0}
            max={12}
            defaultValue={hourState}
            onChange={handleChangeHour}
          >
            <NumberInputField />
            <NumberInputStepper>
              <NumberIncrementStepper />
              <NumberDecrementStepper />
            </NumberInputStepper>
          </NumberInput>
          <NumberInput
            size="xs"
            maxW={16}
            min={0}
            max={60}
            defaultValue={minuteState}
            onChange={handleChangeMinute}
            bg="white"
          >
            <NumberInputField />
            <NumberInputStepper>
              <NumberIncrementStepper />
              <NumberDecrementStepper />
            </NumberInputStepper>
          </NumberInput>
          <NumberInput
            size="xs"
            maxW={16}
            min={0}
            max={60}
            defaultValue={secondState}
            onChange={handleChangeSecond}
            bg="white"
          >
            <NumberInputField />
            <NumberInputStepper>
              <NumberIncrementStepper />
              <NumberDecrementStepper />
            </NumberInputStepper>
          </NumberInput>
        </HStack>

        {/* <EditorContent className="editor-content" editor={editor} /> */}
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
                      borderBottom={"2px"}
                      borderColor={"gray.300"}
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
                        <Button
                          variant="success"
                          w="20%"
                          alignSelf={"flex-end"}
                          onClick={() => {
                            console.log("CLICK")
                            const newComments = allComments.filter(
                              (c) => c.jsonComments !== comment.jsonComments
                            );
                            setAllComments(newComments);
                          }}
                        >
                          Resolve
                        </Button>
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

export default EditableTimeTiptap;

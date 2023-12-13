import { DataItem } from "../../models/qa";
import {
  Box,
  VStack,
  Text,
  AspectRatio,
  Flex,
  IconButton,
  HStack,
  Highlight,
  Tooltip,
} from "@chakra-ui/react";
import Answer from "./Answer.tsx";
import { timeToSeconds } from "../../functions";
import { RepeatIcon } from "@chakra-ui/icons";
import { useEffect, useState } from "react";

interface QAProps {
    data: DataItem;
    query: string;
    tokens: string[];
}
/**
 * Renders a QA video component with the given data, query, and tokens.
 *
 * @param {QAProps} data - The data object containing the video information.
 * @param {string} query - The query string used for highlighting the question.
 * @param {string[]} tokens - The tokens used for highlighting the question.
 * @return {JSX.Element} The rendered QA video component.
 */
function QA_Vdo({ data, query, tokens }: QAProps) {
  const [isQueryTheQuestion, SetisQueryTheQuestion] = useState(false);
  const startTime = timeToSeconds(data.startTime);
  const endTime = timeToSeconds(data.endTime);
  const youtubeURL = `https://www.youtube.com/embed/${data.youtubeURL}?start=${startTime}&end=${endTime}`;
  
  useEffect(() => {
    if (query == data.question) {
      // For Highlighting the question 
      SetisQueryTheQuestion(true);
    }
  }, [query, data.question]);
  
  /**
   * Replays the video by updating the source of the iframe element.
   *
   * @param {string} data.question - The id of the iframe element.
   * @return {void} This function does not return anything.
   */
  const replay = () => {
    const iframe = document.getElementById(data.question) as HTMLImageElement;
    if (iframe) {
      iframe.src = youtubeURL;
    }
  };

  return (
    <Flex
      h="auto"
      gap={4}
      w="full"
      direction={{ base: "column-reverse", md: "column-reverse", lg: "row" }}
    >
      <Box w={{ base: "100%", lg: "65%" }}>
        <VStack spacing={1} alignItems="flex-start">
          <HStack spacing={2} alignItems="center">
            <Text as="b" color="blue">
              เวลาเริ่มต้น {data.startTime} เวลาสิ้นสุด {data.endTime}
            </Text>
            <Tooltip
              hasArrow
              label="กดเพื่อเล่นวิดีโออีกครั้ง"
              bg="gray.300"
              color="black"
              placement="right"
            >
              <IconButton
                aria-label="Play Again"
                icon={<RepeatIcon />}
                onClick={replay}
              />
            </Tooltip>
          </HStack>

          {isQueryTheQuestion == false ? (
            <Text as="b">
              <Highlight query={tokens} styles={{ color: "red" }}>
                {data.question}
              </Highlight>
            </Text>
          ) : (
            <Text as="b" color="red">
              {data.question}
            </Text>
          )}

          <Answer text={data.answer} tokens={tokens} />
        </VStack>
      </Box>
      <Box w={{ base: "100%", lg: "35%" }}>
        <AspectRatio maxW={["560px"]} maxH="300px" ratio={1}>
          <iframe
            id={data.question}
            title={data.question}
            src={youtubeURL}
            allowFullScreen
          />
        </AspectRatio>
      </Box>
    </Flex>
  );
}

export default QA_Vdo;

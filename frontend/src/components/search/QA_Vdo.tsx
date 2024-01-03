import { DataItem } from "../../models/qa";
import { Box, VStack, Text, Flex, Highlight } from "@chakra-ui/react";
import Answer from "./Answer.tsx";
import YoutubeVideo from "./YoutubeVideo.tsx";
import { useEffect, useState } from "react";
import { getCookie } from "typescript-cookie";
import { useNavigate } from "react-router-dom";
import { useRef } from "react";
import { VdoRef } from "./YoutubeVideo.tsx";
import TimesAndTools from "./TimesAndTools.tsx";
import React from "react";
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
  const token = getCookie("token"); //TODO : check token is valid
  const navigate = useNavigate();
  const vdoRef = useRef<VdoRef | null>(null);

  useEffect(() => {
    if (query == data.question) {
      // For Highlighting the question
      SetisQueryTheQuestion(true);
    }
  }, [query, data.question]);

  const handleReplay = () => {
    vdoRef.current?.replay();
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
          <TimesAndTools
            data={data}
            token={token}
            handleReplay={handleReplay}
            navigate={navigate}
          />
          {/* //TODO : Check new component */}
          {/* <HStack spacing={2} alignItems="center">
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
                onClick={handleReplay}
              />
            </Tooltip>
            {token && (
              <Tooltip
                hasArrow
                label="กดเพื่อเสนอข้อแก้ไข"
                bg="gray.300"
                color="black"
                placement="right"
              >
                <IconButton
                  aria-label="Edit"
                  icon={<EditIcon />}
                  onClick={() =>
                    navigate(`/contributor/edit-record/${data.index}`)
                  }
                />
              </Tooltip>
            )}
          </HStack> */}

          {isQueryTheQuestion == false ? (
            <Text variant="question">
              <Highlight query={tokens} styles={{ color: "red" }}>
                {data.question}
              </Highlight>
            </Text>
          ) : (
            <Text variant="question" color="red">
              {data.question}
            </Text>
          )}

          <Answer text={data.answer} tokens={tokens} />
        </VStack>
      </Box>
      <Box w={{ base: "100%", lg: "35%" }}>
        <YoutubeVideo
          ref={vdoRef}
          youtubeURL={data.youtubeURL}
          question={data.question}
          startTime={data.startTime}
          endTime={data.endTime}
        />
      </Box>
    </Flex>
  );
}

export default QA_Vdo;

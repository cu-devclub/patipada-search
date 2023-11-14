import { useState, useEffect, useRef } from "react";
import { Box, IconButton, Highlight } from "@chakra-ui/react";
import { TriangleDownIcon, TriangleUpIcon } from "@chakra-ui/icons";
import useMaxChars from "../hook/useMaxChar.ts"; // Import the custom hook

interface AnswerProps {
  text: string;
  query: string[];
}

function Answer({ text, query }: AnswerProps) {
  const [isExpanded, setIsExpanded] = useState(false);
  const [truncatedText, setTruncatedText] = useState<string | null>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const maxChars = useMaxChars(); // Use the custom hook to calculate maxLines

  useEffect(() => {
    if (containerRef.current) {
      if (text.length > maxChars) {
        setTruncatedText(text.slice(0, maxChars) + "...");
      } else {
        setTruncatedText(null);
      }
    }
  }, [text, maxChars]);

  const toggleReadMore = () => {
    setIsExpanded(!isExpanded);
  };

  return (
    <Box w="full" ref={containerRef} overflow="hidden">
      {truncatedText ? (
        <>
          {isExpanded ? (
            <>
              <Highlight query={query} styles={{ color: "red" }}>
                {text}
              </Highlight>
              <IconButton
                size="xs"
                aria-label="read less"
                onClick={toggleReadMore}
                icon={<TriangleUpIcon w="1rem" h="1rem" />}
              />
            </>
          ) : (
            <>
              <Highlight query={query} styles={{ color: "red" }}>
                {truncatedText}
              </Highlight>
              <IconButton
                size="xs"
                aria-label="read more"
                onClick={toggleReadMore}
                icon={<TriangleDownIcon w="1rem" h="1rem" />}
              />
            </>
          )}
        </>
      ) : (
        <Highlight query={query} styles={{ color: "red" }}>
          {text}
        </Highlight>
      )}
    </Box>
  );
}

export default Answer;

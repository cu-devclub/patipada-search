import { useState, useEffect, useRef } from "react";
import { Box, IconButton, Highlight, Text } from "@chakra-ui/react";
import { TriangleDownIcon, TriangleUpIcon } from "@chakra-ui/icons";
import { useMaxChars } from "../../hook";

/**
 * Renders the Answer component, which displays a text with optional highlighting and a "Read More" or "Read Less" button.
 *
 * @param {AnswerProps} props - The properties for the Answer component.
 *   @param {string} props.text - The text to be displayed.
 *   @param {string[]} props.tokens - The tokens to be highlighted.
 * @return {JSX.Element} The rendered Answer component.
 */

interface AnswerProps {
  text: string;
  tokens: string[];
}
function Answer({ text, tokens }: AnswerProps) {
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
              <Text variant="answer">
                <Highlight query={tokens} styles={{ color: "red" }}>
                  {text}
                </Highlight>{" "}
              </Text>
              <IconButton
                size="xs"
                aria-label="read less"
                onClick={toggleReadMore}
                icon={<TriangleUpIcon w="1rem" h="1rem" />}
              />
            </>
          ) : (
            <>
              <Text variant="answer">
                <Highlight query={tokens} styles={{ color: "red" }}>
                  {truncatedText}
                </Highlight>
              </Text>

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
        <Text variant="answer">
          <Highlight query={tokens} styles={{ color: "red" }}>
            {text}
          </Highlight>
        </Text>
      )}
    </Box>
  );
}

export default Answer;

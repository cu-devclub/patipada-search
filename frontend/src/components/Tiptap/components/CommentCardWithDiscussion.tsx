/* eslint-disable @typescript-eslint/no-explicit-any */

import { Flex, HStack, Button, Box, Text } from "@chakra-ui/react";
import { formatDate } from "../../../functions";
import { ReactNode } from "react";
interface CommentCardProps {
  allComments: any[];
  activeCommentInstance: any;
  commentText: string;
  setCommentText: (comment: string) => void;
  setComment: () => void;
  inputField: ReactNode;
}

function CommentCardWithDiscussion({
  allComments,
  setCommentText,
  activeCommentInstance,
  setComment,
  inputField,
}: CommentCardProps) {
  return (
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
                    borderBottom="2px"
                    borderColor="gray.300"
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
                    </Flex>
                  </Box>
                );
              }
            )}

            {comment.jsonComments.uuid === activeCommentInstance.uuid && (
              <Flex w="full" direction="column" gap={1}>
                {inputField}

                <HStack>
                  <Button onClick={() => setCommentText("")} variant="cancel">
                    Clear
                  </Button>
                  <Button onClick={() => setComment()} variant="success">
                    Add
                  </Button>
                </HStack>
              </Flex>
            )}
          </Box>
        );
      })}
    </Flex>
  );
}

export default CommentCardWithDiscussion;

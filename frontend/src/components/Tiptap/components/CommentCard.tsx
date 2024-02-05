/* eslint-disable @typescript-eslint/no-explicit-any */

import { Flex, HStack, Button, Box, Text } from "@chakra-ui/react";
import { formatDate } from "../../../functions";

interface CommentCardProps {
  allComments: any[];
  resolveComment: (comment: any,time: string) => void;
}

function CommentCard({ allComments, resolveComment }: CommentCardProps) {
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
                        onClick={() =>
                          resolveComment && resolveComment(comment,jsonComment.time)
                        }
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
  );
}

export default CommentCard;

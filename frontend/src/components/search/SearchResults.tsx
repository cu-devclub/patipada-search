import { Flex, VStack, Divider } from "@chakra-ui/react";
import QA_Vdo from "./QA_Vdo.tsx";
import { SearchResultInterface } from "../../models/qa";
/**
 * Renders the list of search results component.
 * The structure is
 * Search Result list
 *  |- QA_Vdo (Question, Answer, Video)
 *    |- Answer (text and `Read more & Read Less` button)
 *
 * @param {SearchResultInterface} data - The search result data.
 * @param {string} query - The search query.
 * @param {Array<string>} tokens - The search query tokens.
 * @return {JSX.Element} The rendered search results component.
 */
function SearchResults({ data, query, tokens }: SearchResultInterface) {
  return (
    <Flex
      flex={1}
      w="full"
      justifyContent={{ base: "center", md: "flex-start" }}
      px={8}
    >
      <VStack spacing={8} w={{ base: "100%", md: "80%", xl: "70%" }}>
        {data.map((item) => (
          <>
            <QA_Vdo data={item} query={query} tokens={tokens} />
            <Divider />
          </>
        ))}
      </VStack>
    </Flex>
  );
}

export default SearchResults;

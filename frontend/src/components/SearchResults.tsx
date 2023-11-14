import { Flex, VStack, Divider } from "@chakra-ui/react";
import QA_Vdo from "./QA_Vdo.tsx";
import { SearchResultsProps } from "../data/dataInterface.tsx";

function SearchResults({ data,query, tokens }: SearchResultsProps) {
  return (
    <Flex
      flex={1}
      w="full"
      justifyContent={{ base: "center", lg: "flex-start" }}
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

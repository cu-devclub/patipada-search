import { Flex, Center, VStack } from "@chakra-ui/react";
import { SearchField } from "../../components/search";
import { Header, Footer, Logo } from "../../components";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

/**
 * Renders a search page with a search field.
 *
 * @returns {React.ReactNode} The rendered search page.
 */
const SearchPage = () => {
  const navigate = useNavigate();
  const [searchParam, SetSearchParam] = useState("");

  /**
   * Sets the search parameters.
   *
   * @param {string} searchParameter - The search parameter to be set.
   * @return {void} This function does not return any value.
   */
  const setSearchParams = (searchParameter: string) => {
    SetSearchParam(searchParameter);
  };

  /**
   * Performs a search using the given search parameter.
   *
   * @param {string} searchParameter - The search parameter to use for the search.
   * @return {void} This function does not return anything.
   */
  const performSearch = (searchParameter: string) => {
    SetSearchParam(searchParameter);
    navigate(`?search=${searchParameter}`);
  };

  return (
    <Flex
      justify="space-between"
      direction="column"
      align="center"
      w="full"
      minH="100svh"
    >
      <VStack w="full" spacing={[4, 16]}>
        <Header />

        <VStack w="full">
          <Center>
            <Logo size={["6xs", "4xs"]} />
          </Center>
          <SearchField
            searchParam={searchParam}
            setSearchParams={setSearchParams}
            performSearch={performSearch}
          />
        </VStack>
      </VStack>
      <Flex w="100%" h="8xs">
        <Footer />
      </Flex>
    </Flex>
  );
};

export default SearchPage;

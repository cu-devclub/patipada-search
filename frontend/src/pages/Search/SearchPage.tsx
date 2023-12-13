import { Flex, Show, Hide, Image } from "@chakra-ui/react";
import { SearchField } from "../../components/search";
import { Header, Footer } from "../../components";
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
      direction="column"
      gap={4}
      justify={"space-between"}
      align="center"
      w="full"
      minH="100svh"
    >
      <Hide below="sm">
        <Header />
      </Hide>

      <Flex
        w="full"
        h="full"
        justify="center"
        direction="column"
        align="center"
        gap="4"
      >
        <Show below="sm">
          <Image
            boxSize="7xs"
            objectFit="cover"
            src="/Dhammanava.svg"
            alt="Dhammanava"
          />
        </Show>

        <SearchField
          searchParam={searchParam}
          setSearchParams={setSearchParams}
          performSearch={performSearch}
        />
      </Flex>
      <Footer />
    </Flex>
  );
}

export default SearchPage;

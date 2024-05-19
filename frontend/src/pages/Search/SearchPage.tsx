import {  VStack, AspectRatio, Grid, GridItem } from "@chakra-ui/react";
import { SearchField } from "../../components/search";
import { Header, Footer } from "../../components";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { WordCloud } from "../../components/logo";

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
    <Grid
      templateRows="0.2fr 2fr 0.2fr"
      templateAreas={`" header"
                        " main"
                        " footer"`}
      gap={4}
      w="full"
      h="100svh"
    >
      <GridItem pl="2" area={"header"}>
        <Header />
      </GridItem>
      <GridItem pl="2" area={"main"}>
        <VStack w="full" spacing={8}>
          <AspectRatio w={{ base: "60%", md: "50%", lg: "30%" }} ratio={16 / 9}>
            <WordCloud />
          </AspectRatio>

          <SearchField
            searchParam={searchParam}
            setSearchParams={setSearchParams}
            performSearch={performSearch}
          />
        </VStack>
      </GridItem>
      <GridItem bg="red" area={"footer"} h="full">
        <Footer />
      </GridItem>
    </Grid>
  );
};

export default SearchPage;

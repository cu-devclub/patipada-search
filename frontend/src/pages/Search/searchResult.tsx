import { SearchResults, HeaderSearch } from "../../components/search";
import { Flex, Divider } from "@chakra-ui/react";
import { Footer } from "../../components";
import { SearchResultInterface, DataItem } from "../../models/qa";
import { useSearchParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import Pagination from "@choc-ui/paginator";

/**
 * Render the search result page.
 *
 * @return {JSX.Element} The JSX element representing the search result page.
 */
function SearchResultPage() {
  const navigate = useNavigate();
  const [queryMessage, SetQueryMessage] = useState("");
  const [data, SetData] = useState<DataItem[]>([]);
  const [tokens, SetTokens] = useState<string[]>([]);

  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 8; // Set the number of items per page here
  const [searchParams] = useSearchParams();
  const query = searchParams.get("search");

  // When query change; get the result from session storage 
  useEffect(() => {
    if (query) {
      SetQueryMessage(query);
      setCurrentPage(1);
      const responseData = sessionStorage.getItem("response");
      if (responseData != null) {
        const { data, tokens }: Pick<SearchResultInterface, "data" | "tokens"> =
          JSON.parse(responseData);
        SetData(data);
        SetTokens(tokens);
      }
    }
  }, [query]);

  // ------ Header Search in result page ---------
  const SetSearchParams = (searchParameter: string) => {
    SetQueryMessage(searchParameter);
  };

  const performSearch = (searchParameter: string) => {
    navigate(`?search=${searchParameter}`);
    location.reload();
  };
  // --------------------------------------------

  // ------- Pagination  ----------------------
  const changePage = (current: number | undefined) => {
    if (current) {
      setCurrentPage(current);
    }
  };

  // Calculate the start and end index for the current page
  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;

  // Get the data for the current page
  let currentPageData = data;
  if (data != null) {
    currentPageData = data.slice(startIndex, endIndex);
  }
  // --------------------------------------------

  return (
    <Flex
      direction="column"
      gap={8}
      justify="space-between"
      align="flex-start"
      w="full"
      minH="100svh"
    >
      {query && (
        <HeaderSearch
          query={query}
          searchParam={queryMessage}
          setSearchParams={SetSearchParams}
          performSearch={performSearch}
        />
      )}
      <Divider />
      {data != null && (
        <>
          <SearchResults
            data={currentPageData}
            query={queryMessage}
            tokens={tokens}
          />
          <Flex w={{ base: "100%", md: "80%", xl: "70%" }} justify={"center"}>
            <Pagination
              current={currentPage}
              total={data.length}
              pageSize={itemsPerPage}
              onChange={(current) => changePage(current)}
              paginationProps={{
                display: "flex",
              }}
              activeStyles={{
                color: "black",
                bg: "blackAlpha.200",
              }}
              hoverStyles={{
                bg: "gray.300",
              }}
            />
          </Flex>
        </>
      )}
      <Footer />
    </Flex>
  );
}

export default SearchResultPage;

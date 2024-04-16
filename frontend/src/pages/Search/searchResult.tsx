import { SearchResults } from "../../components/search";
import { Flex, Grid, GridItem } from "@chakra-ui/react";
import { Footer, HeaderSearch } from "../../components/layout";
import { SearchResultInterface, DataItem } from "../../models/qa";
import { useSearchParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import Pagination from "@choc-ui/paginator";
import {
  SEARCH_STATUS,
  SEARCH_TYPE,
  SearchResultItemsPerPage,
  ToastStatus,
} from "../../constant";
import { searchService } from "../../service/search";
import { MessageToast } from "../../components";

/**
 * Render the search result page.
 *
 * @return {JSX.Element} The JSX element representing the search result page.
 */
function SearchResultPage() {
  const navigate = useNavigate();
  const { addToast } = MessageToast();

  const [queryMessage, SetQueryMessage] = useState("");
  const [data, SetData] = useState<DataItem[]>([]);
  const [tokens, SetTokens] = useState<string[]>([]);
  const [pageNums, SetPageNums] = useState(1);
  const [currentPage, setCurrentPage] = useState(1);
  const [searchParams] = useSearchParams();
  const query = searchParams.get("search");

  // When query change; get the result from session storage
  useEffect(() => {
    if (query) {
      SetQueryMessage(query);
      setCurrentPage(1);
      const responseData = sessionStorage.getItem("response");
      if (responseData != null) {
        const {
          data,
          tokens,
          numPages,
        }: Pick<SearchResultInterface, "data" | "tokens" | "numPages"> =
          JSON.parse(responseData);
        SetData(data);
        SetTokens(tokens);
        SetPageNums(numPages);
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
  const changePage = async (current: number | undefined) => {
    if (current) {
      setCurrentPage(current);
      // fetch search service with new page
      const offset = (current - 1) * SearchResultItemsPerPage;
      await searchService(
        query || "",
        SEARCH_TYPE.DEFAULT,
        SEARCH_STATUS.CONFIRM,
        offset,
        SearchResultItemsPerPage,
        false,
        pageNums
      )
        .then((response: SearchResultInterface) => {
          SetData(response.data);
          window.scrollTo(0, 0);
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    }
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
        {query && (
          <HeaderSearch
            query={query}
            searchParam={queryMessage}
            setSearchParams={SetSearchParams}
            performSearch={performSearch}
          />
        )}
      </GridItem>
      <GridItem pl="2" area={"main"}>
        {data != null && (
          <>
            <SearchResults data={data} query={queryMessage} tokens={tokens} />
            <Flex
              w={{ base: "100%", md: "80%", xl: "70%" }}
              justifyContent={"center"}
              alignItems={"center"}
              pt={4}
            >
              <Pagination
                current={currentPage}
                pageSize={SearchResultItemsPerPage}
                total={pageNums * SearchResultItemsPerPage}
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
                pageNeighbours={2}
              />
            </Flex>
          </>
        )}
      </GridItem>
      <GridItem area={"footer"} h="8xs">
        <Footer />
      </GridItem>
    </Grid>
  );
}

export default SearchResultPage;

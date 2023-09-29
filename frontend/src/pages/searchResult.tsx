import SearchResults from "../components/SearchResults.tsx";
import HeaderSearch from "../components/HeaderSearch.tsx";
import { Flex } from "@chakra-ui/react";
import { Divider } from "@chakra-ui/react";
import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import Pagination from "@choc-ui/paginator";
function SearchResultPage() {
  const navigate = useNavigate();
  const [queryMessage, SetQueryMessage] = useState("");
  const [data, SetData] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 8; // Set the number of items per page here
  const { query } = useParams();
  useEffect(() => {
    if (query) {
      SetQueryMessage(query);
      setCurrentPage(1);
      const responseData = localStorage.getItem("response");
      if (responseData != null) {
        SetData(JSON.parse(responseData));
      }
    }
  }, [query]);

  const SetSearchParams = (searchParameter: string) => {
    SetQueryMessage(searchParameter);
  };

  const performSearch = (searchParameter: string) => {
    navigate(`/result/${searchParameter}`);
    location.reload()
  };
  const changePage = (current: number | undefined) => {
    if (current) {
      setCurrentPage(current);
    }
  };

  // Calculate the start and end index for the current page
  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;

  // Get the data for the current page
  let currentPageData = data
  if (data != null) {
     currentPageData  = data.slice(startIndex, endIndex);
  }

  return (
    <Flex
      direction="column"
      gap={8}
      align="flex-start"
      w="full"
      minH="100vh"
      p={8}
    >
      <HeaderSearch
        searchParam={queryMessage}
        setSearchParams={SetSearchParams}
        performSearch={performSearch}
      />
      <Divider />
      {data != null && (
        <>
          <SearchResults data={currentPageData} query={queryMessage} />
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
    </Flex>
  );
}

export default SearchResultPage;

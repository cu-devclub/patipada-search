import {
  FormControl,
  InputLeftElement,
  InputGroup,
  Flex,
  Text,
  Spinner,
  Tooltip,
} from "@chakra-ui/react";
import {
  AutoComplete,
  AutoCompleteInput,
  AutoCompleteItem,
  AutoCompleteList,
} from "@choc-ui/chakra-autocomplete";
import { SearchIcon } from "@chakra-ui/icons";
import { useState } from "react";
import { searchService } from "../../service/search";
import { SearchResultInterface } from "../../models/qa";
import { SEARCH_STATUS, SEARCH_TYPE, SearchResultItemsPerPage } from "../../constant";
import { MessageToast } from "../toast";
interface SearchOptions {
  key: string;
  question: string;
}

/**
 * Filters the results based on the given term.
 *
 * @param {any} term - The term to filter the results.
 * @return {Promise<SearchOptions[]>} The filtered results.
 */
async function filterResults(term: string) {
  let data: SearchOptions[] = [];
  try {
    const response = await searchService(term);
    if (response) {
      data = response.data.map((item) => ({
        key: item.index,
        question: item.question,
      }));
    }
    return data;
  } catch (error) {
    console.error("Error:", error);
    return data;
  }
}

/**
 * Render a search field component. also controlling the search options
 *
 * @param {SearchFieldProps} {
 *   searchParam, // The current search parameter
 *   setSearchParams, // A function to set the search parameter
 *   performSearch, // A function to perform the search
 * } - The props for the SearchField component
 * @return {JSX.Element} - The rendered search field component
 */

interface SearchFieldProps {
  searchParam: string | null; // Define the searchParam prop
  setSearchParams: (searchParameter: string) => void;
  performSearch: (searchParameter: string) => void;
  offset?: number;
  amount?: number;
}
function SearchField({
  searchParam,
  setSearchParams,
  performSearch,
  offset = 0,
  amount = SearchResultItemsPerPage,
}: SearchFieldProps) {
  const { addToast } = MessageToast();
  const [isLoading, setIsLoading] = useState(false);
  const [options, setOptions] = useState<SearchOptions[]>();
  const [debounceTimer, setDebounceTimer] = useState<NodeJS.Timeout | null>(
    null
  );

  /**
   * Handles the change event of the input.
   * Add debounce timer to prevent multiple API calls
   *
   * @param {Event} evt - The event object.
   * @return {void} This function does not return anything.
   */
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const onChangeInputHandler = (evt: any) => {
    const inputValue = evt.target.value;
    setSearchParams(inputValue);

    // Clear the previous timer if it exists
    if (debounceTimer !== null) {
      clearTimeout(debounceTimer);
    }

    // Set a new timer to wait for a timeout before making the API call
    const timerId = setTimeout(async () => {
      setIsLoading(true);
      const results = await filterResults(inputValue);
      setOptions(results);
      setIsLoading(false);
    }, 500);

    // Save the timer ID for cleanup
    setDebounceTimer(timerId);
  };

  /**
   * Handles the selection of an input.
   *
   * @param {Event} evt - The event object representing the input selection.
   * @return {Promise<void>} A promise that resolves when the function completes.
   */
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  async function onSelectInputHandle(evt: any) {
    let query = evt.item.value;

    // Check if the query is an option key (user selected from options)
    // Have to do this because bug of choc-ui package
    // Where if the options have the same value it mark as the same key
    // and then when user selected, evt.item.value is the index(unable to read)
    // and then when search it will return empty
    const q = options?.find((o) => o.key === query);
    if (q) {
      query = q.question;
    }

    await searchService(
      query,
      SEARCH_TYPE.DEFAULT,
      SEARCH_STATUS.CONFIRM,
      offset,
      amount
    )
      .then((response) => {
        const tokens = [query, ...response.tokens];
        const searchResults: SearchResultInterface = {
          data: response.data,
          query: query,
          tokens: tokens,
          numPages: response.numPages,
        };

        sessionStorage.setItem("response", JSON.stringify(searchResults));

        performSearch(query);
      })
      .catch(() => {
        addToast({
          description: "เกิดข้อผิดพลาดขณะทำการค้นหา",
          status: "error",
        });
      });
  }

  return (
    <FormControl w={{ base: "70%", lg: "50%" }} fontWeight="light">
      <AutoComplete
        emptyState={<Text textAlign="center">ค้นหาเลย</Text>}
        openOnFocus
        isLoading={isLoading}
        onSelectOption={onSelectInputHandle}
        disableFilter
      >
        <InputGroup>
          <InputLeftElement pointerEvents="none" h={["50", "70"]}>
            <SearchIcon color="gray.500" boxSize={6} />
          </InputLeftElement>
          <AutoCompleteInput
            onChange={onChangeInputHandler}
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            onKeyDown={(e: any) => {
              if (e.key === "Enter") {
                onSelectInputHandle(e);
              }
            }}
            variant="search_bar"
            value={searchParam || ""}
            placeholder="ค้นหาเลย"
            h={["50", "70"]}
            fontSize={["md", "lg", "xl"]}
          />
        </InputGroup>
        <AutoCompleteList
          loadingState={
            <div>
              <Spinner
                thickness="4px"
                speed="0.65s"
                emptyColor="gray.200"
                color="blue.500"
                size="md"
              />
            </div>
          }
        >
          {searchParam && (
            <AutoCompleteItem
              key={`self-search`}
              value={searchParam}
              textTransform="capitalize"
              h={["50", "70", "90"]}
              fontSize={["md", "lg", "xl"]}
              w={{ base: "80%", lg: "90%" }}
            >
              <Flex alignItems="center">
                <SearchIcon color="gray.500" boxSize={6} mr={4} />
                <Tooltip
                  hasArrow
                  label={searchParam}
                  bg="gray.300"
                  color="black"
                  placement="right"
                >
                  <Text noOfLines={1}> {searchParam}</Text>
                </Tooltip>
              </Flex>
            </AutoCompleteItem>
          )}
          {options &&
            options.map((obj, idx) => (
              <>
                <AutoCompleteItem
                  key={idx}
                  value={obj.key}
                  textTransform="capitalize"
                  h={["50", "70", "90"]}
                  fontSize={["md", "lg", "xl"]}
                  w="80%"
                >
                  <Flex alignItems="center">
                    <SearchIcon color="gray.500" boxSize={6} mr={4} />
                    <Tooltip
                      hasArrow
                      label={obj.question}
                      bg="gray.300"
                      color="black"
                      placement="right"
                    >
                      <Text noOfLines={1}> {obj.question}</Text>
                    </Tooltip>
                  </Flex>
                </AutoCompleteItem>
              </>
            ))}
        </AutoCompleteList>
      </AutoComplete>
    </FormControl>
  );
}

export default SearchField;

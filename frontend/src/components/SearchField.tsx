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
import axios from "axios";
interface SearchFieldProps {
  searchParam: string | null; // Define the searchParam prop
  setSearchParams: (searchParameter: string) => void;
  performSearch: (searchParameter: string) => void;
}

interface SearchOptions {
  key: string;
  question: string;
}

async function filterResults(term) {
  let data: SearchOptions[] = [];
  try {
    const response = await fetchingData(term);
    // create variable where data type is SearchOptions
    if (response.results != null) {
      data = response.results.map((item) => ({
        key: item.id,
        question: item.question,
      }));
    }
    return data;
  } catch (error) {
    console.error("Error:", error);
    return data;
  }
}

async function fetchingData(query: string) {
  try {
    // const path = "http://localhost:8081"; //* For local development
    const path = import.meta.env.VITE_SEARCH_API_URL; //* For production
    const response = await axios.get(path + `/search?query=${query}`);
    return response.data;
  } catch (error) {
    console.error("Error:", error);
    return [];
  }
}

function SearchField({
  searchParam,
  setSearchParams,
  performSearch,
}: SearchFieldProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [options, setOptions] = useState<SearchOptions[]>();
  const [debounceTimer, setDebounceTimer] = useState<NodeJS.Timeout | null>(
    null
  );

  const onChangeInputHandler = (evt) => {
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
    }, 500); // Adjust the timeout duration as needed (e.g., 500 milliseconds)

    // Save the timer ID for cleanup
    setDebounceTimer(timerId);
  };

  async function onSelectInputHandle(evt) {
    let query = evt.item.value;
    let response = await fetchingData(query);
    if (!response.result && options) {
      const q = options.find((o) => o.key === query);
      if (q) {
        response = await fetchingData(q?.question);
        query = q?.question;
      }
    }
    sessionStorage.setItem("response", JSON.stringify(response.results));

    const tokens = [query, ...response.tokens];
    sessionStorage.setItem("tokens", JSON.stringify(tokens));

    setSearchParams(query);
    performSearch(query);
  }

  return (
    <FormControl w={["90%", "70%", "50%"]}>
      <AutoComplete
        emptyState={<Text textAlign="center">ค้นหาเลย</Text>}
        openOnFocus
        isLoading={isLoading}
        onSelectOption={onSelectInputHandle}
        disableFilter
      >
        <InputGroup>
          <InputLeftElement pointerEvents="none" h={["50", "70", "90"]}>
            <SearchIcon color="gray.500" boxSize={6} />
          </InputLeftElement>
          <AutoCompleteInput
            loadingIcon={
              <div>
                <br />
                <br />
                <Spinner
                  thickness="4px"
                  speed="0.65s"
                  emptyColor="gray.200"
                  color="blue.500"
                  size="md"
                />
              </div>
            }
            onChange={onChangeInputHandler}
            bg="blackAlpha.200"
            pl={12}
            variant="filled"
            value={searchParam}
            placeholder="ค้นหาด้วยคีย์เวิร์ด"
            borderRadius="30"
            h={["50", "70", "90"]}
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
            options.map((obj,idx) => (
              <>
                <AutoCompleteItem
                  key={idx}
                  value={obj.key}
                  textTransform="capitalize"
                  h={["50", "70", "90"]}
                  fontSize={["md", "lg", "xl"]}
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

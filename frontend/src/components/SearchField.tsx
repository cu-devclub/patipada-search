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

const questions = [];

interface SearchFieldProps {
  searchParam: string | null; // Define the searchParam prop
  setSearchParams: (searchParameter: string) => void;
  performSearch: (searchParameter: string) => void;
}

async function filterResults(term) {
  try {
    const data = await fetchingData(term);
    if (data == null) {
      return [];
    }
    const questionsArray = data.map((item) => item.question);

    return questionsArray;
  } catch (error) {
    console.error("Error:", error);
    return [];
  }
}

async function fetchingData(query: string) {
  try {
    // const path = "http://localhost:8081";
    const path = import.meta.env.VITE_SEARCH_API_URL;
    const response = await axios.get(path + `/search?query=${query}`);
    localStorage.setItem("response", JSON.stringify(response.data.results));
    return response.data.results;
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
  const [options, setOptions] = useState(questions);

  const onChangeInputHandler = (evt) => {
    setIsLoading(true);
    setSearchParams(evt.target.value);
    filterResults(evt.target.value).then((results) => {
      setOptions(results);
      setIsLoading(false);
    });
  };

  const onSelectInputHandle = (evt) => {
    setSearchParams(evt.item.value);
    performSearch(evt.item.value);
  };

  return (
    <FormControl w={["90%", "70%", "50%"]}>
      <AutoComplete
        emptyState={<Text textAlign="center">ค้นหาเลย</Text>}
        openOnFocus
        isLoading={isLoading}
        onSelectOption={onSelectInputHandle}
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
          {options.map((question, cid) => (
            <AutoCompleteItem
              key={`option-${cid}`}
              value={question}
              textTransform="capitalize"
              h={["50", "70", "90"]}
              fontSize={["md", "lg", "xl"]}
            >
              <Flex alignItems="center">
                <SearchIcon color="gray.500" boxSize={6} mr={4} />
                <Tooltip
                  hasArrow
                  label={question}
                  bg="gray.300"
                  color="black"
                  placement="right"
                >
                  <Text noOfLines={1}> {question}</Text>
                </Tooltip>
              </Flex>
            </AutoCompleteItem>
          ))}
        </AutoCompleteList>
      </AutoComplete>
    </FormControl>
  );
}

export default SearchField;

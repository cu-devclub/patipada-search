import SearchField from "./SearchField.tsx";
import {
  Flex,
  HStack,
  Center,
  Hide,
  Show,
  IconButton,
  VStack,
  Grid,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { SignInButton } from "../user";
import Logo from "../Logo.tsx";
import { HamburgerIcon } from "@chakra-ui/icons";
interface SearchFieldProps {
  query: string | null;
  searchParam: string; // Define the searchParam prop
  setSearchParams: (searchParameter: string) => void;
  performSearch: (searchParameter: string) => void;
}

/**
 * Renders a search field in the header with the given props.
 *
 * @param {string} query - The current search query.
 * @param {string} searchParam - The current search parameter.
 * @param {Function} setSearchParams - A function to update the search parameters.
 * @param {Function} performSearch - A function to perform the search.
 * @return {JSX.Element} The rendered search field component.
 */
function HeaderSearch({
  query,
  searchParam,
  setSearchParams,
  performSearch,
}: SearchFieldProps) {
  const navigate = useNavigate();
  const [searchParams, SetSearchParams] = useState(
    searchParam !== "" ? searchParam : query
  );
  const ChangeSearch = (s: string) => {
    SetSearchParams(s);
    setSearchParams(s);
  };

  useEffect(() => {
    SetSearchParams(searchParam);
  }, [searchParam]);

  return (
    <Flex
      px={4}
      pt={4}
      direction="row"
      justify="space-between"
      alignItems="center"
      w="full"
    >
      {/* Desktop */}
      <Hide below="md">
        <HStack w="full" gap={4}>
          <Center w="7%" h="7%" onClick={() => navigate("/")}>
            <Logo size={{ md: "6xs" }} />
          </Center>

          <SearchField
            searchParam={searchParams}
            setSearchParams={ChangeSearch}
            performSearch={performSearch}
          />
        </HStack>
        <SignInButton />
      </Hide>
      {/* ------------------------------- */}
      {/* Mobile */}
      <Show below="md">
        <VStack w="full" gap={0}>
          <Grid
            templateColumns="repeat(3, 1fr)"
            gap={6}
            h="7xs"
            alignItems="center"
          >
            <IconButton
              w="10xs"
              aria-label="Open Menu"
              icon={<HamburgerIcon />}
            />
            <Center
              w="full"
              h="full"
              cursor="pointer"
              onClick={() => navigate("/")}
            >
              <Logo size={{ base: "8xs" }} />
            </Center>
            <SignInButton />
          </Grid>
          <SearchField
            searchParam={searchParams}
            setSearchParams={ChangeSearch}
            performSearch={performSearch}
          />
        </VStack>
      </Show>
    </Flex>
  );
}

export default HeaderSearch;

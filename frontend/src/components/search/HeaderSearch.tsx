import SearchField from "./SearchField.tsx";
import {BaseHeader} from "../../components";
import { useState, useEffect } from "react";
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
    <BaseHeader>
      <SearchField
        searchParam={searchParams}
        setSearchParams={ChangeSearch}
        performSearch={performSearch}
      />
    </BaseHeader>
    // <Flex
    //   px={4}
    //   pt={4}
    //   direction="row"
    //   justify="space-between"
    //   alignItems="center"
    //   w="full"
    // >
    //   {/* Desktop */}
    //   <Hide below="md">
    //     <HStack w="full" gap={4}>
    //       <Center w="7%" h="7%" onClick={() => navigate("/")}>
    //         <Logo size={{ md: "6xs" }} />
    //       </Center>

    // <SearchField
    //   searchParam={searchParams}
    //   setSearchParams={ChangeSearch}
    //   performSearch={performSearch}
    // />
    //     </HStack>
    //     {token && username ? (
    //       <UserAvatar username={username} />
    //     ) : (
    //       <SignInButton />
    //     )}
    //   </Hide>
    //   {/* ------------------------------- */}
    //   {/* Mobile */}
    //   <Show below="md">
    //     <VStack w="full" gap={0}>
    //       <Grid
    //         templateColumns="repeat(3, 1fr)"
    //         gap={6}
    //         h="7xs"
    //         alignItems="center"
    //         w="full"
    //       >
    //         <IconButton
    //           w="10xs"
    //           aria-label="Open Menu"
    //           icon={<HamburgerIcon />}
    //         />
    //         <Center
    //           w="full"
    //           h="full"
    //           cursor="pointer"
    //           onClick={() => navigate("/")}
    //         >
    //           <Logo size={{ base: "8xs" }} />
    //         </Center>
    //         {token && username ? (
    //           <Flex justify={"flex-end"}>
    //             <UserAvatar username={username} />
    //           </Flex>
    //         ) : (
    //           <Flex justify={"flex-end"}>
    //             <SignInButton />
    //           </Flex>
    //         )}
    //       </Grid>
    //       <SearchField
    //         searchParam={searchParams}
    //         setSearchParams={ChangeSearch}
    //         performSearch={performSearch}
    //       />
    //     </VStack>
    //   </Show>
    // </Flex>
  );
}

export default HeaderSearch;

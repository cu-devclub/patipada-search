import { Flex, Show, Hide, Image } from "@chakra-ui/react";
import SearchField from "../components/SearchField.tsx";
import Header from "../components/Header.tsx";
import Footer from "../components/Footer.tsx";
// import { QuestionOutlineIcon } from "@chakra-ui/icons";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
function SearchPage() {
  const navigate = useNavigate();
  const [searchParam, SetSearchParam] = useState("");
  const setSearchParams = (searchParameter: string) => {
    SetSearchParam(searchParameter);
  };
  const performSearch = (searchParameter: string) => {
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

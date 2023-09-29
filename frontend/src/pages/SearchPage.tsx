import { Flex, Show, Hide ,Image} from "@chakra-ui/react";
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
    navigate(`/result/${searchParameter}`);
  };
  return (
    <Flex
      direction="column"
      gap={4}
      justify={{ base: "center", sm: "space-between" }}
      align="center"
      w="full"
      minH="100vh"
    >
      <Hide below="sm">
        <Header />
      </Hide>

      <Show below="sm">
        <Image
          boxSize="7xs"
          objectFit="cover"
          src="Logo.svg"
          alt="Dhammanava"
        />
      </Show>

      <SearchField
        searchParam={searchParam}
        setSearchParams={setSearchParams}
        performSearch={performSearch}
      />

      <Hide below="sm">
        <Footer />
      </Hide>
    </Flex>
  );
}

export default SearchPage;

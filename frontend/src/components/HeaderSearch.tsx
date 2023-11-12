import SearchField from "./SearchField.tsx";
import { Flex, Image } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
interface SearchFieldProps {
  query: string | null;
  searchParam: string; // Define the searchParam prop
  setSearchParams: (searchParameter: string) => void;
  performSearch: (searchParameter: string) => void;
}

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
    <Flex direction="row" gap={8} alignItems="center" w="full">
      <Image
        boxSize="7xs"
        src="/Dhammanava.svg"
        alt="Dhammanava"
        borderRadius="full"
        onClick={() => navigate("/")}
        cursor={"pointer"}
      />

      <SearchField
        searchParam={searchParams}
        setSearchParams={ChangeSearch}
        performSearch={performSearch}
      />
    </Flex>
  );
}

export default HeaderSearch;

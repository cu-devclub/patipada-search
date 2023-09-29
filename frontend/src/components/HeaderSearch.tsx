import SearchField from "./SearchField.tsx";
import { Flex,Image } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
interface SearchFieldProps {
  searchParam: string | null; // Define the searchParam prop
  setSearchParams: (searchParameter: string) => void;
  performSearch: (searchParameter: string) => void;
}

function HeaderSearch({
  searchParam,
  setSearchParams,
  performSearch,
}: SearchFieldProps) {
  const navigate = useNavigate();
  return (
    <Flex direction="row" gap={8} alignItems="center" w="full">
      <Image
        boxSize="7xs"
        src="/Logo.svg"
        alt="Dhammanava"
        borderRadius="full"
        onClick={()=>navigate("/")}
        cursor={"pointer"}
      />

      <SearchField
        searchParam={searchParam}
        setSearchParams={setSearchParams}
        performSearch={performSearch}
      />
    </Flex>
  );
}

export default HeaderSearch;

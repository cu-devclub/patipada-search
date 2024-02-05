import SearchField from "../../search/SearchField.tsx";
import BaseHeader  from "./BaseHeader.tsx";
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
  );
}

export default HeaderSearch;

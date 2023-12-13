import { useSearchParams } from "react-router-dom";
import SearchPage from "./SearchPage";
import SearchResultPage from "./searchResult";
/**
 * Executes the search middleware.
 *
 * @return {JSX.Element} The rendered search result page or search page.
 * depends on the search parameter in the URL
*/
function SearchMiddleware() {
  const [searchParams] = useSearchParams();

  const hasParam = searchParams.get("search");

  if (hasParam) {
    return <SearchResultPage />;
  }
  return <SearchPage />;
}

export default SearchMiddleware;

import { useSearchParams } from 'react-router-dom';
import SearchPage from "./SearchPage";
import SearchResultPage from "./searchResult";
function Wrapper() {
  const [searchParams] = useSearchParams();

  const hasParam = searchParams.get("search");

  if (hasParam) {
    // only A and not B
    return <SearchResultPage />;
  } 
  return <SearchPage/>; // render nothing
}

export default Wrapper
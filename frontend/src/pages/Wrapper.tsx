import { useSearchParams } from 'react-router-dom';
import SearchPage from "./SearchPage";
import SearchResultPage from "./searchResult";
function Wrapper() {
  const [searchParams] = useSearchParams();

  const hasParam = searchParams.get("search");

  if (hasParam) {
    return <SearchResultPage />;
  } 
  return <SearchPage/>; 
}

export default Wrapper
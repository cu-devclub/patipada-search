import { BrowserRouter, Routes, Route } from "react-router-dom";
import SearchPage from "./pages/SearchPage";
import SearchResultPage from "./pages/searchResult";
export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<SearchPage />} />
        <Route path="/result">
          <Route index element={<SearchPage />} />
          <Route path=":query" element={<SearchResultPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

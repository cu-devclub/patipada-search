import React from "react";
import ReactDOM from "react-dom/client";
// import SearchPage from './SearchPage.tsx'
// import SearchResultPage from './SearchResultPage.tsx'
import { ChakraProvider } from "@chakra-ui/react";
import theme from "./theme/index.ts";
import "../src/theme/styles.css";
import App from "./App.tsx";
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ChakraProvider theme={theme}>
      <App />
    </ChakraProvider>
  </React.StrictMode>
);

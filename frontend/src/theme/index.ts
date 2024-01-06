import { extendTheme } from "@chakra-ui/react";
import component from "./component";
const font = {
  fonts: {
    heading: `'Mitr', sans-serif`,
    body: `'Mitr', sans-serif`,
    mono: `Mitr`,
  },
  fontSizes: {
    "15xl" : "15rem",
  }
};

const size = {
  sizes: {
    "4xs": "10rem",
    "5xs": "8rem",
    "6xs": "7rem",
    "7xs": "5rem",
    "8xs": "4rem",
    "10xs": "2rem",
  },
};
const color = {
  colors: {
    brand_orange: {
      400: "#FFA800",
    },
    gray: {
      450: "#D9D9D9",
    },
    brand_gray : {
      400: "#F2F2F2",
      500: "#D9D9D9",
    },
  },
};

const theme = extendTheme(font, component, size, color);
export default theme;

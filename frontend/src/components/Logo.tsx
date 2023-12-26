import logo from "/Dhammanava.svg";
import { Image } from "@chakra-ui/react";
const Logo = ({size}) => {
  return (
    <Image
      boxSize={size}
      src={logo}
      alt="Dhammanava"
      objectFit="contain"
    />
  );
};

export default Logo;

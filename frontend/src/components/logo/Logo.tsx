import logo from "/Dhammanava.svg";
import { Image } from "@chakra-ui/react";
interface Props {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  size: any;
}
const Logo = ({ size }:Props) => {
  return (
    <Image boxSize={size} src={logo} alt="Dhammanava" objectFit="contain" />
  );
};

export default Logo;

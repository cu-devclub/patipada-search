import { Flex } from "@chakra-ui/react";
import { SignInButton } from "./user";
function Header() {
  return (
    <Flex pr={4} justify="flex-end" alignItems="center" w="100%" h="8xs">
      <SignInButton />
    </Flex>
  );
}

export default Header;

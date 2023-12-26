import { Flex } from "@chakra-ui/react";
import { SignInButton,UserAvatar } from "./user";
import { getCookie } from "typescript-cookie";
function Header() {
  const token = getCookie("token");
  const username = getCookie("username");
  return (
    <Flex pr={4} justify="flex-end" alignItems="center" w="100%" h="8xs">
      {token&&username ? <UserAvatar username={username}/> : <SignInButton />}
    </Flex>
  );
}

export default Header;

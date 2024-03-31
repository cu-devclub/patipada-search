import { Center, Flex } from "@chakra-ui/react";
import { SignInButton, UserAvatar } from "../../user";
import { getCookie } from "typescript-cookie";
import { Logo } from "../..";
function Header() {
  const token = getCookie("token");
  const username = getCookie("username");
  return (
    <Flex
      pr={4}
      justify="space-between"
      alignItems="center"
      w="100%"
      h="6xs"
      px={2}
    >
      <Center w="7%" h="7%">
        <Logo size={{ md: "6xs" }} />
      </Center>

      {token && username ? (
        <UserAvatar username={username} />
      ) : (
        <SignInButton />
      )}
    </Flex>
  );
}

export default Header;

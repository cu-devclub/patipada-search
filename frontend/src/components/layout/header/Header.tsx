import { Box, Flex } from "@chakra-ui/react";
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
      // alignItems="center"
      w="100%"
      h="6xs"
      px={2}
    >
      <Box w={{ base: "20%", lg: "20%" }} h="full">
        <Logo size={{base:"8xs", lg: "7xs" }} />
      </Box>

      {token && username ? (
        <UserAvatar username={username} />
      ) : (
        <SignInButton marginTop="4" />
      )}
    </Flex>
  );
}

export default Header;

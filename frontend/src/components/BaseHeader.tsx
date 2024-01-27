import {
  Flex,
  HStack,
  Center,
  Hide,
  Show,
  IconButton,
  VStack,
  Grid,
  Divider,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { SignInButton, UserAvatar } from "./user/index.ts";
import Logo from "./Logo.tsx";
import { HamburgerIcon } from "@chakra-ui/icons";
import { getCookie } from "typescript-cookie";

interface BaseHeaderProps {
  children: React.ReactNode;
}

function BaseHeader({ children }: BaseHeaderProps) {
  const navigate = useNavigate();
  const token = getCookie("token");
  const username = getCookie("username");

  return (
    <VStack >
      <Flex
        px={4}
        pt={4}
        direction="row"
        justify="space-between"
        alignItems="center"
        w="full"
      >
        {/* Desktop */}
        <Hide below="md">
          <HStack w="full" gap={4}>
            <Center w="7%" h="7%" onClick={() => navigate("/")}>
              <Logo size={{ md: "6xs" }} />
            </Center>

            {children}
          </HStack>
          {token && username ? (
            <UserAvatar username={username} />
          ) : (
            <SignInButton />
          )}
        </Hide>
        {/* ------------------------------- */}
        {/* Mobile */}
        <Show below="md">
          <VStack w="full" gap={0}>
            <Grid
              templateColumns="repeat(3, 1fr)"
              gap={6}
              h="7xs"
              alignItems="center"
              w="full"
            >
              <IconButton
                w="10xs"
                aria-label="Open Menu"
                icon={<HamburgerIcon />}
              />
              <Center
                w="full"
                h="full"
                cursor="pointer"
                onClick={() => navigate("/")}
              >
                <Logo size={{ base: "8xs" }} />
              </Center>
              {token && username ? (
                <Flex justify={"flex-end"}>
                  <UserAvatar username={username} />
                </Flex>
              ) : (
                <Flex justify={"flex-end"}>
                  <SignInButton />
                </Flex>
              )}
            </Grid>
            {children}
          </VStack>
        </Show>
      </Flex>
      <Divider />
    </VStack>
  );
}

export default BaseHeader;

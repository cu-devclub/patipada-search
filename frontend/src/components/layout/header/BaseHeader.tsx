import {
  Flex,
  HStack,
  Center,
  Hide,
  Show,
  VStack,
  Divider,
  Popover,
  PopoverArrow,
  PopoverBody,
  PopoverCloseButton,
  PopoverContent,
  PopoverHeader,
  PopoverTrigger,
  IconButton,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { SignInButton, UserAvatar } from "../../user/index.ts";
import Logo from "../../logo/Logo.tsx";
import { getCookie } from "typescript-cookie";
import { InfoIcon } from "@chakra-ui/icons";
interface BaseHeaderProps {
  children?: React.ReactNode;
  showDivider?: boolean;
}

const LoginInfo = () => {
  return (
    <Popover>
      <PopoverTrigger>
        <IconButton
          icon={<InfoIcon />}
          aria-label="Login Information"
          variant="ghost"
        />
      </PopoverTrigger>
      <PopoverContent>
        <PopoverArrow />
        <PopoverCloseButton />
        <PopoverHeader>ลงชื่อเข้าใช้</PopoverHeader>
        <PopoverBody>ลงชื่อเข้าใช้สำหรับ Content Contributor</PopoverBody>
      </PopoverContent>
    </Popover>
  );
};

function BaseHeader({ children, showDivider = false }: BaseHeaderProps) {
  const navigate = useNavigate();
  const token = getCookie("token");
  const username = getCookie("username");

  return (
    <VStack>
      <Flex
        px={2}
        pt={2}
        direction="row"
        justify="space-between"
        alignItems="center"
        w="full"
      >
        {/* Desktop */}
        <Hide below="md">
          <HStack w="full" gap={4}>
            <Center
              cursor="pointer"
              w="7%"
              h="7%"
              onClick={() => navigate("/")}
            >
              <Logo size={{ md: "6xs" }} />
            </Center>

            {children}
          </HStack>
          {token && username ? (
            <UserAvatar username={username} />
          ) : (
            <Flex alignItems={"center"}>
              <SignInButton />
              <LoginInfo />
            </Flex>
          )}
        </Hide>
        {/* ------------------------------- */}
        {/* Mobile */}
        <Show below="md">
          <HStack w="full" gap={4}>
            <Center
              cursor="pointer"
              w="20%"
              h="20%"
              onClick={() => navigate("/")}
            >
              <Logo size={{ md: "6xs" }} />
            </Center>
            {children}
          </HStack>
          {token && username ? (
            <Flex>
              <UserAvatar username={username} />
            </Flex>
          ) : (
            <SignInButton />
          )}
        </Show>
      </Flex>
      {showDivider && <Divider />}
    </VStack>
  );
}

export default BaseHeader;

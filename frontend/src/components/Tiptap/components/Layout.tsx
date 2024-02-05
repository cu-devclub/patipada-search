import { Flex } from "@chakra-ui/react";
import { ReactNode } from "react";

interface LayoutProps {
  children: ReactNode;
}
export const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <Flex
      dir="row"
      w="100%"
      bg="gray.100"
      p={2}
      borderRadius={"lg"}
      position="relative"
    >
      {children}
    </Flex>
  );
};

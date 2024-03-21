import { Box, Heading, VStack } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
interface AdminSidebarProps {
  activePage: string;
}
function AdminSidebar({ activePage }: AdminSidebarProps) {
  const navigate = useNavigate();

  const pages = ["Dashboard", "Request", "Users"];

  return (
    <Box w="full" h="full" bg="brand_orange.400" pl={2}>
      <VStack align="start" pt={8}>
        {pages.map((page,key) => (
          <Heading
            key={key}
            fontWeight={page === activePage ? "bold" : "normal"}
            onClick={() => navigate(`/admin/${page.toLowerCase()}`)}
            cursor="pointer"
          >
            {page}
          </Heading>
        ))}
      </VStack>
    </Box>
  );
}

export default AdminSidebar;

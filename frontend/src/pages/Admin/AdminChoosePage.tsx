import { Button, Flex, Heading, Stack, Text } from "@chakra-ui/react";
import { useLocation, useNavigate } from "react-router-dom";
function AdminChoosePage() {
  const navigate = useNavigate();
  const location = useLocation();
  
  return (
    <Flex
      w="full"
      minH="100svh"
      justify="center"
      align="center"
      direction="column"
      gap={4}
    >
      <Heading>เลือกเว็บไซต์ที่ต้องการเข้าใช้งาน</Heading>
      <Stack direction={["column", "row"]}>
        <Button variant="brand" onClick={() =>{
          if (location.state?.from) {
            navigate(location.state.from);
          }
          else {
            navigate("/");
          }
        }}>
          <Text fontWeight={"normal"} fontSize={"24"} p={4}>
            ค้นหาข้อมูล
          </Text>
        </Button>
        <Button variant="brand" onClick={() => navigate("/admin/dashboard")}>
          <Text fontWeight={"normal"} fontSize={"24"} p={4}>
            Admin Dashboard
          </Text>
        </Button>
      </Stack>
    </Flex>
  );
}

export default AdminChoosePage;

import { AuthenForm } from "../../components/admin";
import { Center } from "@chakra-ui/react";
function AdminLoginPage() {
  return (
    <Center w="100%" h="100vh">
      <AuthenForm title="Admin Login" isShowForgetPassword={true} />
    </Center>
  );
}

export default AdminLoginPage;

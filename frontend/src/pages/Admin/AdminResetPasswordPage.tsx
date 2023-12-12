import { AuthenForm } from "../../components/admin";
import { Center } from "@chakra-ui/react";
function AdminResetPasswordPage() {
  return (
    <Center w="100%" h="100vh">
      <AuthenForm title="Reset Password" isShowForgetPassword={false} />
    </Center>
  );
}

export default AdminResetPasswordPage;

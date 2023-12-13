import { ForgetPasswordForm } from "../../components/admin";
import { Center } from "@chakra-ui/react";
function AdminResetPasswordPage() {
  const submit = (email) => {
    console.log(email);
    //TODO : making request to backend
  }
  return (
    <Center w="100%" h="100vh">
      <ForgetPasswordForm submit={submit}/>
    </Center>
  );
}

export default AdminResetPasswordPage;

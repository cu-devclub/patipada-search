import { AuthenForm } from "../../components/admin";
import { MessageToast } from "../../components";
import { Center } from "@chakra-ui/react";
import {login} from "../../service/admin";
import {ToastStatus} from "../../constant"
import { setCookie } from "typescript-cookie";
import { useNavigate } from "react-router-dom";
function AdminLoginPage() {
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const submit = (username: string, password: string) => {
    login(username, password)
      .then((response) => {
        addToast({
          description: "Login successfully",
          status: ToastStatus.SUCCESS,
        });
        setCookie("token", response.token);
        navigate("/admin")
      })
      .catch((error) => {
        if(error.status) {
          if (error.status === 401){
            addToast({
              description: "Incorrect username or password",
              status: ToastStatus.WARNING,
            });
          }
        } else {
          addToast({
            description: "Login failed",
            status: ToastStatus.ERROR,
          });
        }
      })
  };
  return (
    <Center w="100%" h="100vh">
      <AuthenForm
        submit={submit}
      />
    </Center>
  );
}

export default AdminLoginPage;

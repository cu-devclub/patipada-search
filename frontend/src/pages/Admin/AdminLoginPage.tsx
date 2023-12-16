import { AuthenForm } from "../../components/admin";
import { MessageToast } from "../../components";
import { Center } from "@chakra-ui/react";
import { login } from "../../service/admin";
import { ToastStatus } from "../../constant";
import { setCookie } from "typescript-cookie";
import { useNavigate } from "react-router-dom";

/**
 * Admin Login Page Component.
 *
 * Renders a login form and handles the submission of login requests.
 */
const AdminLoginPage = () =>  {
  const navigate = useNavigate();
  const { addToast } = MessageToast();

  /**
   * Submits a login request using the given username and password.
   *
   * @param {string} username - The username for the login request.
   * @param {string} password - The password for the login request.
   */
  const submit = (username: string, password: string) => {
    login(username, password)
      .then((response) => {
        addToast({
          description: "Login successfully",
          status: ToastStatus.SUCCESS,
        });
        setCookie("token", response.token);
        navigate("/admin");
      })
      .catch((error) => {
        if (error.status) {
          if (error.status === 401) {
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
      });
  };
  
  return (
    <Center w="100%" h="100vh">
      <AuthenForm submit={submit} />
    </Center>
  );
}

export default AdminLoginPage;

import { AuthenForm } from "../../components/user/forms";
import { MessageToast, Logo } from "../../components";
import { Flex, Heading, VStack, Text, HStack, Button } from "@chakra-ui/react";
import { login } from "../../service/user";
import { ToastStatus } from "../../constant";
import { setCookie } from "typescript-cookie";
import { useNavigate } from "react-router-dom";
import { LoginDTO } from "../../models/user";
import { useState } from "react";
import { ReturnError } from "../../service/error";
/**
 * Admin Login Page Component.
 *
 * Renders a login form and handles the submission of login requests.
 */
const LoginPage = () => {
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const [formError, setformError] = useState(false);

  /**
   * Submits a login request using the given username and password.
   *
   * @param {string} username - The username for the login request.
   * @param {string} password - The password for the login request.
   */
  const submit = async (username: string, password: string) => {
    const loginDTO: LoginDTO = {
      username: username,
      password: password,
    };

    await login(loginDTO)
      .then((response) => {
        setformError(false);
        addToast({
          description: "Login successfully",
          status: ToastStatus.SUCCESS,
        });
        setCookie("token", response.token);
        setCookie("username", username);
        setCookie("role", response.role);
        navigate(-1);
        // if (response.role == Role.ADMIN || response.role == Role.SUPER_ADMIN) navigate("/user");
        // else navigate("/")
      })
      .catch((error: ReturnError) => {
        setformError(true);
        addToast({
          description: error.message,
          status: error.toastStatus,
        });
      });
  };

  return (
    <Flex
      w="100%"
      minH="100svh"
      bg="gray.600"
      justify={"flex-start"}
      align={"center"}
      direction={"column"}
      pt={12}
    >
      <Logo size="7xs" />
      <VStack spacing={0} pb={4}>
        <Heading
          fontSize={"5xl"}
          color={"whiteAlpha.900"}
          letterSpacing={"tighter"}
          textShadow={"0px 4px 4px rgba(0, 0, 0, 0.25)"}
        >
          ลงชื่อเข้าใช้งาน
        </Heading>
        <HStack spacing={2}>
          <Text color={"whiteAlpha.900"} fontSize={"lg"}>
            ยังไม่มีบัญชีใช่ไหม ?
          </Text>
          <Button
            variant="brand_link"
            fontSize={"lg"}
            onClick={() => navigate("/user/register")}
          >
            สมัครเลย
          </Button>
        </HStack>
      </VStack>
      <AuthenForm submit={submit} formError={formError} />
    </Flex>
  );
};

export default LoginPage;

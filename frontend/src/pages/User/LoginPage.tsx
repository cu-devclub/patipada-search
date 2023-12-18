import { AuthenForm } from "../../components/user/forms";
import { MessageToast, Logo } from "../../components";
import { Flex, Heading, VStack, Text, HStack, Button } from "@chakra-ui/react";
import { login } from "../../service/user";
import { ToastStatus } from "../../constant";
import { setCookie } from "typescript-cookie";
import { useNavigate } from "react-router-dom";

/**
 * Admin Login Page Component.
 *
 * Renders a login form and handles the submission of login requests.
 */
const LoginPage = () => {
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
      <AuthenForm submit={submit} />
    </Flex>
  );
};

export default LoginPage;

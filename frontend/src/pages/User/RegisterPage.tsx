import { RegisterForm } from "../../components/user/forms";
import { Logo, MessageToast } from "../../components";
import { Flex, Heading, VStack } from "@chakra-ui/react";
import { useState } from "react";
import {
  Role,
  ToastStatus,
  ERR_Messages_MAP,
  SERVER_ERROR_MESSAGE,
} from "../../constant";
import { RegisterDTO, LoginDTO } from "../../models/user";
import { register } from "../../service/user";
import { login } from "../../service/user";
import { setCookie } from "typescript-cookie";
import { useNavigate } from "react-router-dom";
import { ReturnError } from "../../service/error";

const RegisterPage = () => {
  const { addToast } = MessageToast();
  const navigate = useNavigate();

  // 400 Bad request has 2 different message
  // 1. `Username already exists`
  // 2. `Email already exists`
  // const [errMessage, setErrMessage] = useState("");
  const [usernameError, setusernameError] = useState(false);
  const [emailError, setemailError] = useState(false);

  const submit = async (username: string, email: string, password: string) => {
    const registerDTO: RegisterDTO = {
      username: username,
      email: email,
      password: password,
      role: Role.USER,
    };

    await register(registerDTO)
      .then(() => {
        addToast({
          description: "register successfully",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch((error:ReturnError) => {
         if (error.status === 400) {
           if (error.message === SERVER_ERROR_MESSAGE.USERNAME_ALREADY_EXISTS) {
             setusernameError(true);
           } else if (error.message === SERVER_ERROR_MESSAGE.EMAIL_ALREADY_EXISTS) {
             setemailError(true);
           }
           addToast({
             description: ERR_Messages_MAP[error.message],
             status: ToastStatus.WARNING,
           });
         } else {
            addToast({
              description: error.message,
              status: error.toastStatus,
            });
         }
      });

    afterwardsLogin(username, password);
  };

  const afterwardsLogin = async (username: string, password: string) => {
    const loginDTO: LoginDTO = {
      username: username,
      password: password,
    };

    await login(loginDTO)
      .then((response) => {
        addToast({
          description: "Login successfully",
          status: ToastStatus.SUCCESS,
        });
        setCookie("token", response.token);
      })
      .catch((error : ReturnError) => {
        addToast({
          description: error.message,
          status: error.toastStatus,
        });
      });

      navigate("/user");

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
          ลงทะเบียน
        </Heading>
      </VStack>
      <RegisterForm
        submit={submit}
        usernameError={usernameError}
        emailError={emailError}
      />
    </Flex>
  );
};

export default RegisterPage;

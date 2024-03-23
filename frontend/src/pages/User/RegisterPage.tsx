import { RegisterForm } from "../../components/user/forms";
import { MessageToast } from "../../components";
import { Button, Flex, Heading } from "@chakra-ui/react";
import { useState } from "react";
import {
  Role,
  ToastStatus,
  ERR_Messages_MAP,
  SERVER_ERROR_MESSAGE,
} from "../../constant";
import { RegisterDTO, LoginDTO } from "../../models/user";
import { registerService, loginService } from "../../service/user";
import { setCookie } from "typescript-cookie";
import { useNavigate, useLocation } from "react-router-dom";
import { ReturnError } from "../../service/error";
import { UserBasePage } from "./UserBasePage";
const RegisterPage = () => {
  const { addToast } = MessageToast();
  const navigate = useNavigate();
  const location = useLocation();

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

    await registerService(registerDTO)
      .then(() => {
        addToast({
          description: "register successfully",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch((error: ReturnError) => {
        if (error.status === 400) {
          if (error.message === SERVER_ERROR_MESSAGE.USERNAME_ALREADY_EXISTS) {
            setusernameError(true);
          } else if (
            error.message === SERVER_ERROR_MESSAGE.EMAIL_ALREADY_EXISTS
          ) {
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

    await loginService(loginDTO)
      .then((response) => {
        addToast({
          description: "Login successfully",
          status: ToastStatus.SUCCESS,
        });
        setCookie("token", response.token);
        if (response.role == Role.ADMIN || response.role == Role.SUPER_ADMIN) {
          navigate("/admin/choosePage");
        } else {
          if (location.state?.from) {
            navigate(location.state.from);
          } else {
            navigate("/");
          }
        }
      })
      .catch((error: ReturnError) => {
        addToast({
          description: error.message,
          status: error.toastStatus,
        });
      });
  };

  return (
    <UserBasePage>
      <Heading
        fontSize={["3xl", "5xl"]}
        color={"whiteAlpha.900"}
        letterSpacing={"tighter"}
        textShadow={"0px 4px 4px rgba(0, 0, 0, 0.25)"}
        textAlign={"center"}
        pb={2}
      >
        ลงทะเบียนสำหรับ <br />
        Content Contributor
      </Heading>
      <RegisterForm
        submit={submit}
        usernameError={usernameError}
        emailError={emailError}
      />
      <Flex alignSelf={"flex-end"}>
        <Button
          variant="brand_link"
          color="blue.100"
          onClick={() =>
            navigate("/user/login", { state: { from: location.state?.from } })
          }
        >
          กลับหน้าเข้าสู่ระบบ
        </Button>
      </Flex>
    </UserBasePage>
  );
};

export default RegisterPage;

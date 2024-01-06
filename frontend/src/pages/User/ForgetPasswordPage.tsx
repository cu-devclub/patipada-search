import { ForgetPasswordForm } from "../../components/user/forms";
import { Text, Heading, Button, Flex } from "@chakra-ui/react";
import { MessageToast } from "../../components";
import { forgetPassword } from "../../service/user";
import { useState } from "react";
import { ReturnError } from "../../service/error";
import { UserBasePage } from "./UserBasePage";
import { useNavigate, useLocation } from "react-router-dom";
function ForgetPasswordPage() {
  const { addToast } = MessageToast();
  const navigate = useNavigate();
  const location = useLocation();
  const [formSuccess, setformSuccess] = useState(false);
  const [emailError, setemailError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const submit = async (email: string) => {
    setIsLoading(true);
    await forgetPassword(email)
      .then(() => {
        setemailError(false);
        setformSuccess(true);
      })
      .catch((err: ReturnError) => {
        setemailError(true);
        addToast({
          description: err.message,
          status: err.toastStatus,
        });
      });
    setIsLoading(false);
  };

  return (
    <UserBasePage>
      <Heading
        fontSize={"5xl"}
        color={"whiteAlpha.900"}
        letterSpacing={"tighter"}
        textShadow={"0px 4px 4px rgba(0, 0, 0, 0.25)"}
      >
        ลืมรหัสผ่าน
      </Heading>
      <Text color={"whiteAlpha.900"} fontSize={"lg"} pb={4}>
        ลิ้งค์เปลี่ยนรหัสผ่านจะถูกส่งไปยังอีเมลของท่าน
      </Text>
      <ForgetPasswordForm
        submit={submit}
        formSuccess={formSuccess}
        emailError={emailError}
        isLoading={isLoading}
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
}

export default ForgetPasswordPage;

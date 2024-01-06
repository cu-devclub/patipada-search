import { ResetPasswordForm } from "../../components/user/forms";
import {  Text, Heading } from "@chakra-ui/react";
import {  MessageToast } from "../../components";
import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { verifyResetPasswordToken, resetPassword } from "../../service/user";
import { ReturnError } from "../../service/error";
import { UserBasePage } from "./UserBasePage";
function ResetPasswordPage() {
  const { token } = useParams();
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const [formError, setformError] = useState(false);
  useEffect(() => {
    const verifyToken = async (token: string) => {
      await verifyResetPasswordToken(token)
        .then((res) => {
          if (res.result == false) {
            alert("url หมดอายุ หากต้องการเปลี่ยนรหัสผ่าน กรุณาขอ url ใหม่");
            navigate("/user/login");
          }
        })
        .catch((err: ReturnError) => {
          addToast({
            description: err.message,
            status: err.toastStatus,
          });
        });
    };

    if (token) verifyToken(token);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [token]);

  const submit = async (password: string) => {
    if (!token) return;
    await resetPassword(token, password)
      .then(() => {
        setformError(false);
        addToast({
          description: "เปลี่ยนรหัสผ่านสำเร็จ",
          status: "success",
        });
        navigate("/user/login");
      })
      .catch((err: ReturnError) => {
        setformError(true);
        addToast({
          description: err.message,
          status: err.toastStatus,
        });
      });
  };

  return (
    <UserBasePage>
      <Heading
        fontSize={"5xl"}
        color={"whiteAlpha.900"}
        letterSpacing={"tighter"}
        textShadow={"0px 4px 4px rgba(0, 0, 0, 0.25)"}
      >
        เปลี่ยนรหัสผ่าน
      </Heading>
      <Text color={"whiteAlpha.900"} fontSize={"lg"} pb={2}>
        กรุณากรอกรหัสผ่านใหม่
      </Text>
      <ResetPasswordForm submit={submit} formError={formError} />
    </UserBasePage>
  );
}

export default ResetPasswordPage;

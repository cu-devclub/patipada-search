import { ChangePasswordForm } from "../../components/user/forms";
import { Flex, Text, VStack, Heading } from "@chakra-ui/react";
import { Logo, MessageToast } from "../../components";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { changePassword } from "../../service/user";
import { ReturnError } from "../../service/error";
import { getCookie } from "typescript-cookie";

function ChangePasswordPage() {
  const token = getCookie("token") || "";
  const username = getCookie("username") || "";
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const [oldPasswordError, setOldPasswordError] = useState(false); // not equal in db Password
  const [newPasswordError, setNewPasswordError] = useState(false); // == db Password

  const submit = async (oldPassword: string, newPassword: string) => {
    if (username == "" || token == "") return;
    await changePassword(token, oldPassword, newPassword)
      .then(() => {
        addToast({
          description: "เปลี่ยนรหัสผ่านสำเร็จ",
          status: "success",
        });
       navigate(-1)
      })
      .catch((err: ReturnError) => {
        addToast({
          description: err.message,
          status: err.toastStatus,
        });
        if (err.status == 401) {
            setOldPasswordError(true);
        }
        else if (err.status == 422) {
            setNewPasswordError(true);
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
          เปลี่ยนรหัสผ่าน
        </Heading>
        <Text color={"whiteAlpha.900"} fontSize={"lg"}>
          กรุณากรอกรหัสผ่านใหม่
        </Text>
      </VStack>
      <ChangePasswordForm
        username={username}
        submit={submit}
        oldPasswordError={oldPasswordError}
        newPasswordError={newPasswordError}
      />
    </Flex>
  );
}

export default ChangePasswordPage;

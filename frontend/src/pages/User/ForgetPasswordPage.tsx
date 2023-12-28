import { ForgetPasswordForm } from "../../components/user/forms";
import { Flex, Text, VStack, Heading } from "@chakra-ui/react";
import { Logo, MessageToast } from "../../components";
import { forgetPassword } from "../../service/user";
import { useState } from "react";
import { ReturnError } from "../../service/error";

function ForgetPasswordPage() {
  const { addToast } = MessageToast();

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
          ลืมรหัสผ่าน
        </Heading>
        <Text color={"whiteAlpha.900"} fontSize={"lg"}>
          ลิ้งค์เปลี่ยนรหัสผ่านจะถูกส่งไปยังอีเมลของท่าน
        </Text>
      </VStack>
      <ForgetPasswordForm
        submit={submit}
        formSuccess={formSuccess}
        emailError={emailError}
        isLoading={isLoading}
      />
    </Flex>
  );
}

export default ForgetPasswordPage;

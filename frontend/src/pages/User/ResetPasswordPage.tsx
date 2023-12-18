import { ForgetPasswordForm } from "../../components/user/forms";
import { Flex, Text, VStack, Heading } from "@chakra-ui/react";
import { Logo } from "../../components";

function ResetPasswordPage() {
  const submit = (email) => {
    console.log(email);
    //TODO : making request to backend
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
          ลิ้งค์เปลี่ยนรหัสผ่านจะถูกส่งไปยังอีเมลของท่าน
        </Text>
      </VStack>
      <ForgetPasswordForm submit={submit} />
    </Flex>
  );
}

export default ResetPasswordPage;

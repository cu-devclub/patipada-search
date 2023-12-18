import { RegisterForm } from "../../components/user/forms";
import { Logo } from "../../components";
import { Flex, Heading, VStack } from "@chakra-ui/react";

const RegisterPage = () => {
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
      <RegisterForm />
    </Flex>
  );
};

export default RegisterPage;

import { EmailIcon } from "@chakra-ui/icons";
import {
  Box,
  FormControl,
  Flex,
  FormLabel,
  Input,
  Stack,
  Button,
  InputLeftElement,
  InputGroup,
  Text,
  VStack,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { isValidEmail,handleEnterKeyPress } from "../../../functions";
interface FormProps {
  submit: (email: string) => void;
  formSuccess: boolean;
  emailError: boolean;
}

export default function ForgetPasswordForm({
  submit,
  formSuccess,
  emailError,
}: FormProps) {
  const [email, setEmail] = useState("");
  const [submitCount, setsubmitCount] = useState(0);
  const [tempCredential, setTempCredential] = useState({
    email: "",
  });
    const verifyChangeCredential =
      tempCredential.email != email 

  const navigate = useNavigate();
  const iconColor = formSuccess ? "green.400" : "gray.600";
  const emailFieldVariant = formSuccess
    ? `success_authen_field`
    : `authen_field`;
  const isEmailInvalid =
    (submitCount > 0 && !isValidEmail(email)) || (emailError && !verifyChangeCredential);
  const errMessage = emailError ? "ไม่พบ email นี้ในระบบ" : "อีเมลไม่ถูกต้อง";

  const submitForm = () => {
    setsubmitCount(submitCount + 1);
    if (isEmailInvalid || !isValidEmail(email)) return;
    setTempCredential({ email:email });
    submit(email);
  };

  return (
    <Box w={["xs", "md"]}>
      <Box rounded={"lg"} bg={"white"} boxShadow={"lg"} p={8} w="full">
        <Stack spacing={2} w="full">
          <FormControl id="email" isRequired isInvalid={isEmailInvalid}>
            <FormLabel fontWeight={"light"}>อีเมล</FormLabel>
            <InputGroup>
              <InputLeftElement>
                <EmailIcon color={iconColor} />
              </InputLeftElement>
              <Input
                type="email"
                variant={emailFieldVariant}
                onChange={(e) => setEmail(e.target.value)}
                onKeyDown={handleEnterKeyPress(submitForm)}
              />
            </InputGroup>
            <FormErrorMessage>{errMessage}</FormErrorMessage>
          </FormControl>
          <VStack pt={2}>
            <Button variant="brand" onClick={submitForm}>
              ส่งอีเมล
            </Button>
            {formSuccess && (
              <Text color="green.400">
                ระบบได้ส่งลิ้งค์สำหรับเปลี่ยนรหัสผ่านไปที่อีเมลล์ของท่านแล้ว
              </Text>
            )}
          </VStack>
        </Stack>
      </Box>
      <Flex justify="flex-end">
        <Button
          variant="brand_link"
          onClick={() => navigate("/user/login")}
          color="blue.100"
        >
          กลับหน้าเข้าสู่ระบบ
        </Button>
      </Flex>
    </Box>
  );
}

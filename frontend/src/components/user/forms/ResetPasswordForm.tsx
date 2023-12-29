import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import {
  Box,
  FormControl,
  Center,
  FormLabel,
  Input,
  Stack,
  Button,
  InputGroup,
  InputRightElement,
  IconButton,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useState } from "react";
import { isLengthEnough, handleEnterKeyPress } from "../../../functions";
import { PASSWORD_REQUIRED_LENGTH } from "../../../constant";
interface FormProps {
  submit: (password: string) => void;
  formError: boolean;
}

export default function ResetPasswordForm({ submit, formError }: FormProps) {
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [tempCredential, setTempCredential] = useState({
    password: "",
  });

  const verifyChangeCredential = tempCredential.password != password;

  const passwordErrMessage = formError
    ? "เกิดข้อผิดพลาด"
    : "รหัสผ่านต้องมีความยาวมากกว่า 8 ตัวอักษร";
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const [countSubmit, SetCountSubmit] = useState(0);
  const isPasswordInvalid =
    (countSubmit > 0 && !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH)) ||
    (formError && !verifyChangeCredential);
  const isConfirmPasswordInvalid =
    countSubmit > 0 && password !== confirmPassword;

  const submitForm = () => {
    SetCountSubmit(countSubmit + 1);
    if (
      isPasswordInvalid ||
      isConfirmPasswordInvalid ||
      !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH)
    )
      return;
    setTempCredential({ password: password });
    submit(password);
  };

  return (
    <Box w={["xs", "md"]}>
      <Box rounded={"lg"} bg={"white"} boxShadow={"xl"} p={8} w="full">
        <Stack spacing={2} w="full">
          <FormControl id="password" isRequired isInvalid={isPasswordInvalid}>
            <FormLabel fontWeight={"light"}>รหัสผ่าน</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={showPassword ? "text" : "password"}
                onChange={(e) => setPassword(e.target.value)}
                variant={"authen_field"}
                onKeyDown={handleEnterKeyPress(submitForm)}
              />

              <InputRightElement width="3rem">
                <IconButton
                  size="sm"
                  h="1.75rem"
                  aria-label="View/Hide password"
                  onClick={() => setShowPassword(!showPassword)}
                  icon={showPassword ? <ViewIcon /> : <ViewOffIcon />}
                />
              </InputRightElement>
            </InputGroup>
            <FormErrorMessage>{passwordErrMessage}</FormErrorMessage>
          </FormControl>

          <FormControl
            id="confirm-password"
            isRequired
            isInvalid={isConfirmPasswordInvalid}
          >
            <FormLabel fontWeight={"light"}>ยืนยันรหัสผ่าน</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={showConfirmPassword ? "text" : "password"}
                onChange={(e) => setConfirmPassword(e.target.value)}
                variant={"authen_field"}
                onKeyDown={handleEnterKeyPress(submitForm)}
              />

              <InputRightElement width="3rem">
                <IconButton
                  size="sm"
                  h="1.75rem"
                  aria-label="View/Hide password"
                  onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                  icon={showConfirmPassword ? <ViewIcon /> : <ViewOffIcon />}
                />
              </InputRightElement>
            </InputGroup>
            <FormErrorMessage>
              {isConfirmPasswordInvalid
                ? "รหัสผ่านไม่ตรงกัน"
                : "กรุณากรอกรหัสผ่าน"}
            </FormErrorMessage>
          </FormControl>

          <Center pt={2}>
            <Button variant="brand" onClick={submitForm}>
              เปลี่ยนรหัสผ่าน
            </Button>
          </Center>
        </Stack>
      </Box>
    </Box>
  );
}

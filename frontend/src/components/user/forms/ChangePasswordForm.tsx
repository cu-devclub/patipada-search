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
import { isLengthEnough,handleEnterKeyPress } from "../../../functions";
import { PASSWORD_REQUIRED_LENGTH } from "../../../constant";
interface FormProps {
  username: string;
  submit: (oldPassword: string, newPassword: string) => void;
  oldPasswordError: boolean;
  newPasswordError: boolean;
}

export default function ChangePasswordForm({
  username,
  submit,
  oldPasswordError,
  newPasswordError,
}: FormProps) {
  const [oldPassword, setOldPassword] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [tempCredential, setTempCredential] = useState({
    oldPassword: "",
    password: "",
  });

  const verifyChangeCredential =
    tempCredential.oldPassword != oldPassword ||
    tempCredential.password != password;

  const oldPasswordErrMessage = oldPasswordError
    ? "รหัสผ่านเดิมไม่ตรงกับข้อมูลในระบบ"
    : "รหัสผ่านต้องมีความยาวมากกว่า 8 ตัวอักษร";

  const passwordErrMessage = newPasswordError
    ? "รหัสผ่านใหม่เหมือนกับรหัสผ่านเดิมในระบบ"
    : "รหัสผ่านต้องมีความยาวมากกว่า 8 ตัวอักษร";
  const [showOldPassword, setShowOldPassword] = useState(false);

  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const [countSubmit, SetCountSubmit] = useState(0);

  const isOldPasswordInvalid =
    (countSubmit > 0 &&
      !isLengthEnough(oldPassword, PASSWORD_REQUIRED_LENGTH)) ||
    (oldPasswordError && !verifyChangeCredential);

  const isPasswordInvalid =
    (countSubmit > 0 && !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH)) ||
    (newPasswordError && !verifyChangeCredential);

  const isConfirmPasswordInvalid =
    countSubmit > 0 && password !== confirmPassword;

  const submitForm = () => {
    SetCountSubmit(countSubmit + 1);

    if (
      isPasswordInvalid ||
      isConfirmPasswordInvalid ||
      !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH) ||
      password !== confirmPassword
    )
      return;
    setTempCredential({ oldPassword: oldPassword, password: password });
    submit(oldPassword, password);
  };

  return (
    <Box w={["xs", "md"]}>
      <Box rounded={"lg"} bg={"white"} boxShadow={"xl"} p={8} w="full">
        <Stack spacing={2} w="full">
          <FormControl id="username" isDisabled={true}>
            <FormLabel fontWeight={"light"}>ชื่อผู้ใช้งาน</FormLabel>
            <Input
              pr="3rem"
              type={"text"}
              value={username}
              variant={"authen_field"}
              onKeyDown={handleEnterKeyPress(submitForm)}
            />
          </FormControl>

          <FormControl
            id="oldPassword"
            isRequired
            isInvalid={isOldPasswordInvalid}
          >
            <FormLabel fontWeight={"light"}>รหัสผ่านเดิม</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={showOldPassword ? "text" : "password"}
                onChange={(e) => setOldPassword(e.target.value)}
                variant={"authen_field"}
                onKeyDown={handleEnterKeyPress(submitForm)}
              />

              <InputRightElement width="3rem">
                <IconButton
                  size="sm"
                  h="1.75rem"
                  aria-label="View/Hide password"
                  onClick={() => setShowOldPassword(!showOldPassword)}
                  icon={showOldPassword ? <ViewIcon /> : <ViewOffIcon />}
                />
              </InputRightElement>
            </InputGroup>
            <FormErrorMessage>{oldPasswordErrMessage}</FormErrorMessage>
          </FormControl>

          <FormControl id="password" isRequired isInvalid={isPasswordInvalid}>
            <FormLabel fontWeight={"light"}>รหัสผ่านใหม่</FormLabel>
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

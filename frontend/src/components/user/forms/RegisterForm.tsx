import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import {
  Flex,
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
import { useNavigate } from "react-router-dom";
import { isValueExist, isLengthEnough, isValidEmail } from "../../../functions";
import { PASSWORD_REQUIRED_LENGTH } from "../../../constant";

export default function RegisterForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [email, setEmail] = useState("");

  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const [countSubmit, SetCountSubmit] = useState(0);
  const isUsernameInValid = countSubmit > 0 && !isValueExist(username);
  const isEmailInvalid = countSubmit > 0 && !isValidEmail(email);
  const isPasswordInvalid =
    countSubmit > 0 && !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH);
  const isConfirmPasswordInvalid =
    countSubmit > 0 && password !== confirmPassword;

  const navigate = useNavigate();

  const submitForm = () => {
    SetCountSubmit(countSubmit + 1);
    // submit(username, password);
  };

  return (
    <Box w={["xs", "md"]}>
      <Box rounded={"lg"} bg={"white"} boxShadow={"xl"} p={8} w="full">
        <Stack spacing={2} w="full">
          <FormControl id="username" isRequired isInvalid={isUsernameInValid}>
            <FormLabel fontWeight={"light"}>ชื่อผู้ใช้งาน</FormLabel>
            <Input
              type="text"
              onChange={(e) => setUsername(e.target.value)}
              variant={"authen_field"}
            />
            <FormErrorMessage>กรุณากรอกชื่อผู้ใช้งาน</FormErrorMessage>
          </FormControl>

          <FormControl id="email" isRequired isInvalid={isEmailInvalid}>
            <FormLabel fontWeight={"light"}>อีเมล</FormLabel>
            <Input
              type="email"
              onChange={(e) => setEmail(e.target.value)}
              variant={"authen_field"}
            />
            <FormErrorMessage>อีเมลไม่ถูกต้อง</FormErrorMessage>
          </FormControl>

          <FormControl id="password" isRequired isInvalid={isPasswordInvalid}>
            <FormLabel fontWeight={"light"}>รหัสผ่าน</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={showPassword ? "text" : "password"}
                onChange={(e) => setPassword(e.target.value)}
                variant={"authen_field"}
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
            <FormErrorMessage>
              รหัสผ่านต้องมีความยาวมากกว่า 8 ตัวอักษร
            </FormErrorMessage>
          </FormControl>

          <FormControl
            id="confirm-password"
            isRequired
            isInvalid={isConfirmPasswordInvalid || isPasswordInvalid}
          >
            <FormLabel fontWeight={"light"}>ยืนยันรหัสผ่าน</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={showConfirmPassword ? "text" : "password"}
                onChange={() => setConfirmPassword(password)}
                variant={"authen_field"}
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
              ลงทะเบียน
            </Button>
          </Center>
        </Stack>
      </Box>
      <Flex justify="flex-end">
        <Button
          variant="brand_link"
          color="blue.100"
          onClick={() => navigate("/user/login")}
        >
          กลับหน้าเข้าสู่ระบบ
        </Button>
      </Flex>
    </Box>
  );
}

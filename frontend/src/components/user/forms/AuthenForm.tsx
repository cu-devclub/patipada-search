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
import {
  isValueExist,
  isLengthEnough,
  handleEnterKeyPress,
} from "../../../functions";
import { PASSWORD_REQUIRED_LENGTH } from "../../../constant";
interface FormProps {
  submit: (username: string, password: string) => void;
  formError: boolean;
  locationState: string;
}

export default function AuthenForm({ submit, formError, locationState }: FormProps) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [show, setShow] = useState(false);
  const handleClick = () => setShow(!show);
  const navigate = useNavigate();
  const [submitCount, setSubmitCount] = useState(0);
  const [tempCredential, setTempCredential] = useState({
    username: "",
    password: "",
  });
  const userNameErrorMessage = formError
    ? "ชื่อผู้ใช้งานหรือรหัสผ่านผิด"
    : "กรุณากรอกชื่อผู้ใช้งาน";
  const passwordErrorMessage = formError
    ? "ชื่อผู้ใช้งานหรือรหัสผ่านผิด"
    : "รหัสผ่านต้องมีความยาวมากกว่า 8 ตัวอักษร";

  const verifyChangeCredential =
    tempCredential.username != username || tempCredential.password != password;

  const isUsernameInvalid =
    (submitCount > 0 && !isValueExist(username)) ||
    (formError && !verifyChangeCredential);

  const isPasswordInValid =
    (submitCount > 0 &&
      (!isValueExist(password) ||
        !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH))) ||
    (formError && !verifyChangeCredential);

  const submitForm = () => {
    setSubmitCount(submitCount + 1);
    if (
      isUsernameInvalid ||
      isPasswordInValid ||
      !isValueExist(username) ||
      !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH)
    ) {
      return;
    }
    submit(username, password);
    setTempCredential({ username: username, password: password });
  };

  return (
    <Box w={["2xs", "md"]} pb={12}>
      <Box rounded={"lg"} bg={"white"} boxShadow={"xl"} p={8} w="full">
        <Stack spacing={2} w="full">
          <FormControl id="username" isRequired isInvalid={isUsernameInvalid}>
            <FormLabel fontWeight={"light"}>ชื่อผู้ใช้งาน</FormLabel>
            <Input
              type="text"
              onChange={(e) => setUsername(e.target.value)}
              variant={"authen_field"}
              onKeyDown={handleEnterKeyPress(submitForm)}
            />
            <FormErrorMessage>{userNameErrorMessage}</FormErrorMessage>
          </FormControl>
          <FormControl id="password" isRequired isInvalid={isPasswordInValid}>
            <FormLabel fontWeight={"light"}>รหัสผ่าน</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={show ? "text" : "password"}
                onChange={(e) => setPassword(e.target.value)}
                variant={"authen_field"}
                onKeyDown={handleEnterKeyPress(submitForm)}
              />

              <InputRightElement width="3rem">
                <IconButton
                  size="sm"
                  h="1.75rem"
                  aria-label="View/Hide password"
                  onClick={handleClick}
                  icon={show ? <ViewIcon /> : <ViewOffIcon />}
                />
              </InputRightElement>
            </InputGroup>
            <FormErrorMessage>{passwordErrorMessage}</FormErrorMessage>
          </FormControl>
          <Stack spacing={0}>
            <Flex direction="column" align={"end"}>
              <Button
                variant="brand_link"
                onClick={() => navigate("/user/forget-password", {state: { from: locationState }})}
              >
                ลืมรหัสผ่าน
              </Button>
            </Flex>
            <Center>
              <Button variant="brand" onClick={submitForm}>
                ลงชื่อเข้าใช้
              </Button>
            </Center>
          </Stack>
        </Stack>
      </Box>
    </Box>
  );
}

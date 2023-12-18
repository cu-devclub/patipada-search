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
import { isValueExist, isLengthEnough } from "../../../functions";
import { PASSWORD_REQUIRED_LENGTH } from "../../../constant";
interface FormProps {
  submit: (username: string, password: string) => void;
}

export default function AuthenForm({ submit }: FormProps) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [show, setShow] = useState(false);
  const handleClick = () => setShow(!show);
  const navigate = useNavigate();
  const [isUsernameValid, setisUsernameValid] = useState(true);

  const [submitCount, setSubmitCount] = useState(0);
  const isPasswordInValid =
    submitCount > 0 &&
    (!isValueExist(password) ||
      !isLengthEnough(password, PASSWORD_REQUIRED_LENGTH));

  const submitForm = () => {
    setSubmitCount(submitCount + 1);
    if (!isValueExist(username)) {
      setisUsernameValid(false);
      return;
    }
    submit(username, password);
  };

  return (
    <Box w={["xs", "md"]}>
      <Box rounded={"lg"} bg={"white"} boxShadow={"xl"} p={8} w="full">
        <Stack spacing={2} w="full">
          <FormControl id="username" isRequired isInvalid={!isUsernameValid}>
            <FormLabel fontWeight={"light"}>ชื่อผู้ใช้งาน</FormLabel>
            <Input
              type="text"
              onChange={(e) => setUsername(e.target.value)}
              variant={"authen_field"}
            />
            <FormErrorMessage>กรุณากรอกชื่อผู้ใช้งาน</FormErrorMessage>
          </FormControl>
          <FormControl id="password" isRequired isInvalid={isPasswordInValid}>
            <FormLabel fontWeight={"light"}>รหัสผ่าน</FormLabel>
            <InputGroup>
              <Input
                pr="3rem"
                type={show ? "text" : "password"}
                onChange={(e) => setPassword(e.target.value)}
                variant={"authen_field"}
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
            <FormErrorMessage>
              รหัสผ่านต้องมีความยาวมากกว่า 8 ตัวอักษร
            </FormErrorMessage>
          </FormControl>
          <Stack spacing={0}>
            <Flex direction="column" align={"end"}>
              <Button
                variant="brand_link"
                onClick={() => navigate("/user/reset-password")}
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

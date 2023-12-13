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
  Heading,
  useColorModeValue,
  InputGroup,
  InputRightElement,
} from "@chakra-ui/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
interface FormProps {
  submit: (username:string,password:string) => void;
}

export default function AuthenForm({ submit}: FormProps) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [show, setShow] = useState(false);
  const handleClick = () => setShow(!show);
  const navigate = useNavigate();

  const OnChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(e.target.value);
  };

  const OnChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const submitForm = () => {
    submit(username,password);
  };

  return (
    <Box w="50%">
      <Stack spacing={8} mx={"auto"} maxW={"lg"} py={12} px={6}>
        <Stack align={"center"}>
          <Heading fontSize={"4xl"}>Login</Heading>
        </Stack>
        <Box
          rounded={"lg"}
          bg={useColorModeValue("white", "gray.700")}
          boxShadow={"lg"}
          p={8}
          w="full"
        >
          <Stack spacing={4} w="full">
            <FormControl id="username" isRequired>
              <FormLabel>Username</FormLabel>
              <Input type="text" onChange={OnChangeUsername} />
            </FormControl>
            <FormControl id="password" isRequired>
              <FormLabel>Password</FormLabel>
              <InputGroup>
                <Input
                  type={show ? "text" : "password"}
                  onChange={OnChangePassword}
                />
                <InputRightElement width="4.5rem">
                  <Button h="1.75rem" size="sm" onClick={handleClick}>
                    {show ? <ViewIcon /> : <ViewOffIcon />}
                  </Button>
                </InputRightElement>
              </InputGroup>
            </FormControl>
            <Stack spacing={7}>
              <Flex direction="column" align={"end"}>
                <Button
                  colorScheme="blue"
                  variant="link"
                  onClick={() => navigate("/admin-reset-password")}
                >
                  Forget Password?
                </Button>
              </Flex>
              <Center>
                <Button
                  bg={"blue.400"}
                  color={"white"}
                  _hover={{
                    bg: "blue.500",
                  }}
                  onClick={submitForm}
                  w="50%"
                >
                  Login
                </Button>
              </Center>
            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Box>
  );
}

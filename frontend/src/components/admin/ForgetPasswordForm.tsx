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
} from "@chakra-ui/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
interface FormProps {
  submit: (email: string) => void;
}

export default function ForgetPasswordForm({ submit}: FormProps) {
  const [email, setEmail] = useState("");
  
  const navigate = useNavigate();

  const OnChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };



  const submitForm = () => {
    submit(email);
  };

  return (
    <Box w="50%">
      <Stack spacing={8} mx={"auto"} maxW={"lg"} py={12} px={6}>
        <Stack align={"center"}>
          <Heading fontSize={"4xl"}>Change password</Heading>
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
              <FormLabel>Email</FormLabel>
              <Input type="text" onChange={OnChangeEmail} />
            </FormControl>
            
            <Stack spacing={7}>
              <Flex direction="column" align={"end"}>
                <Button
                  colorScheme="blue"
                  variant="link"
                  onClick={() => navigate("/admin")}
                >
                  Back To Login
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
                  Send Email
                </Button>
              </Center>
            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Box>
  );
}

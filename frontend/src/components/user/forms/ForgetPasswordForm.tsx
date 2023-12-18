import { EmailIcon } from "@chakra-ui/icons";
import {
  Box,
  FormControl,
  Center,
  FormLabel,
  Input,
  Stack,
  Button,
  InputLeftElement,
  InputGroup,
  Flex,
  FormErrorMessage,
} from "@chakra-ui/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { isValidEmail } from "../../../functions";
interface FormProps {
  submit: (email: string) => void;
}

export default function ForgetPasswordForm({ submit }: FormProps) {
  const [email, setEmail] = useState("");
  const [submitCount, setsubmitCount] = useState(0)
  const navigate = useNavigate();

  const isEmailInvalid = submitCount > 0 && !isValidEmail(email);

  const submitForm = () => {
    setsubmitCount(submitCount + 1)
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
                <EmailIcon color={"gray.600"} />
              </InputLeftElement>
              <Input
                type="email"
                variant={"authen_field"}
                onChange={(e) => setEmail(e.target.value)}
              />
            </InputGroup>
            <FormErrorMessage>อีเมลไม่ถูกต้อง</FormErrorMessage>
          </FormControl>
          <Center pt={2}>
            <Button variant="brand" onClick={submitForm}>
              ส่งอีเมล
            </Button>
          </Center>
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

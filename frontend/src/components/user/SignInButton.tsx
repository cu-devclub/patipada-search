import { Button } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

interface SignInButtonProps {
  marginTop?: string;
}

function SignInButton({ marginTop }: SignInButtonProps = { marginTop: "0" }) {
  const navigate = useNavigate();
  return (
    <Button
      variant={"brand"}
      fontSize={{ base: "12", lg: "16" }}
      mt={marginTop}
      onClick={() =>
        navigate("/user/login", {
          state: { from: location.pathname + location.search },
        })
      }
    >
      ลงชื่อเข้าใช้
    </Button>
  );
}

export default SignInButton;

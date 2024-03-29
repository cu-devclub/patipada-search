import { Button } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
function SignInButton() {
  const navigate = useNavigate();
  return (
    <Button
      variant={"brand"}
      fontSize={{ base: "12", lg: "16" }}
      onClick={() => navigate('/user/login', { state: { from: location.pathname + location.search} })}
    >
      ลงชื่อเข้าใช้
    </Button>
  );
}

export default SignInButton;

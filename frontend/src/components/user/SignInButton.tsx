import { Button } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
function SignInButton() {
  const navigate = useNavigate();
  return (
    <Button variant={"brand"} fontSize={{base:"12",lg:"16"}} onClick={()=> navigate("/user/login")} >
      ลงชื่อเข้าใช้
    </Button>
  );
}

export default SignInButton
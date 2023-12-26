import { removeCookie } from "typescript-cookie";
// import { useNavigate } from "react-router-dom";
export const SignOut = () => {
  removeCookie("token");
  removeCookie("username");
  window.location.reload()
//   const navigate = useNavigate();
//   navigate(0);
};

import { getCookie, removeCookie } from "typescript-cookie";
import { verifyToken, authorize } from "../service/user";
export const SignOut = () => {
  removeCookie("token");
  removeCookie("username");
  removeCookie("role");

  if (location.pathname.startsWith("/admin")) {
    window.location.href = "/user/login";
  } else {
    window.location.reload();
  }
};

export const ValidateToken = async () => {
  const token = getCookie("token");
  let ch = false;
  if (token) {
    ch = await verifyToken(token)
      .then((res) => {
        if (res == true) {
          return true;
        } else {
          SignOut();
        }
        return false;
      })
      .catch(() => {
        SignOut();
        return false;
      });
  }
  return ch;
};

export const AuthorizeAdmin = async (requireRole: string) => {
  const token = getCookie("token");
  let ch = false;
  if (token) {
    ch = await authorize(token, requireRole)
      .then((res) => {
        if (res == true) {
          return true;
        } else {
          SignOut();
        }
        return false;
      })
      .catch(() => {
        SignOut();
        return false;
      });
  }
  return ch;
};

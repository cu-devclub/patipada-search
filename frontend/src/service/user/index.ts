import {login} from "./login.ts";
import {register} from "./register.ts";
import { forgetPassword } from "./forgetPassword.ts";
import { verifyResetPasswordToken,resetPassword } from "./resetPassword.ts";
import { changePassword } from "./changePassword.ts";
import { verifyToken, authorize } from "./token.ts";
export {
  login,
  register,
  forgetPassword,
  verifyResetPasswordToken,
  resetPassword,
  changePassword,
  verifyToken,
  authorize,
};
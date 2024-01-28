// generate async function register that takes username, password, email, and role as parameters

import axios from '../axiosInstance';
import { CreateCustomError, ReturnError } from "../error";
import { RegisterDTO } from "../../models/user";
import { ERR_Messages, ToastStatus } from "../../constant";
import { authURL } from '../../constant/serviceURL';
export const register = async (registerDTO: RegisterDTO) => {
  try {
    const response = await axios.post(`${authURL}/register`, {
      username: registerDTO.username,
      password: registerDTO.password,
      email: registerDTO.email,
      role: registerDTO.role,
    });
    return response.data;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status == 400) {
      returnError = {
        message: requestError.message,
        status: 400,
        toastStatus: ToastStatus.WARNING,
      };
    } else if (requestError.status === 401) {
      returnError = {
        message: ERR_Messages.NOT_ATTACH_TOKEN,
        status: 401,
        toastStatus: ToastStatus.ERROR,
      };
    } else if (requestError.status === 409) {
      returnError = {
        message: ERR_Messages.NO_PERMISSION_REGISTER,
        status: 401,
        toastStatus: ToastStatus.ERROR,
      };
    } else {
      returnError = {
        message: ERR_Messages.SYSTEM_ERROR,
        status: 500,
        toastStatus: ToastStatus.ERROR,
      };
    }
    throw returnError;
  }
};

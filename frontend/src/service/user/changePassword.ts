import axios from "../axiosInstance";
import { ERR_Messages, ToastStatus } from "../../constant";
import { CreateCustomError, ReturnError } from "../error";
import { authURL } from "../../constant/serviceURL";
import { getCookie } from "typescript-cookie";
export const changePasswordService = async (
  token: string,
  oldPassword: string,
  newPassword: string
) => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie("token");
    const response = await axios.post(
      `${authURL}/change-password`,
      {
        oldPassword: oldPassword,
        newPassword: newPassword,
      },
      {
        headers: {
          Authorization: token,
        },
      }
    );
    return response.data;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status == 400) {
      returnError = {
        message: ERR_Messages.INVALID_PASSWORD_FORMAT,
        status: 400,
        toastStatus: ToastStatus.ERROR,
      };
    } else if (requestError.status === 401) {
      returnError = {
        message: ERR_Messages.INVALID_OLD_PASSWORD,
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

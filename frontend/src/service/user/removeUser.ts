import axios from "../axiosInstance";
import { ERR_Messages, ToastStatus } from "../../constant";
import { CreateCustomError, ReturnError } from "../error";
import { authURL } from "../../constant/serviceURL";
import { getCookie } from "typescript-cookie";

export const removeUser = async (id: string) => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie("token");

    const response = await axios.delete(`${authURL}/user/${id}`);
    return response.data.status;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status == 400) {
      returnError = {
        message: ERR_Messages.INVALID_ID,
        status: 400,
        toastStatus: ToastStatus.ERROR,
      };
    } else if (requestError.status === 401) {
      returnError = {
        message: ERR_Messages.INVALID_TOKEN,
        status: 401,
        toastStatus: ToastStatus.ERROR,
      } 
    } else if (requestError.status === 403) {
      returnError = {
        message: ERR_Messages.NO_PERMISSION_DELETE,
        status: 401,
        toastStatus: ToastStatus.ERROR,
      };
    } else if (requestError.status === 404) {
      returnError = {
        message: ERR_Messages.NOT_FOUND,
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

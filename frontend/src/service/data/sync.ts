import axios from "../axiosInstance";
import { ERR_Messages, ToastStatus } from "../../constant";
import { CreateCustomError, ReturnError } from "../error";
import { getCookie } from "typescript-cookie";
import { dataURL } from "../../constant/serviceURL";

export const syncRequestRecordService = async (requestID: string) => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie("token");
    const response = await axios.post(`${dataURL}/sync`, {
      requestID: requestID,
    });
    return response.data;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status === 400) {
      returnError = {
        message: ERR_Messages.BAD_REQUEST,
        status: 400,
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

export const syncAllRequestRecordService = async () => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie("token");
    const response = await axios.post(`${dataURL}/sync-all`);
    return response.data;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status === 400) {
      returnError = {
        message: ERR_Messages.BAD_REQUEST,
        status: 400,
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

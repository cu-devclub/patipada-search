import axios from '../axiosInstance';
import { CreateCustomError, ReturnError } from '../error';
import { ToastStatus, ERR_Messages } from '../../constant';
import { authURL } from '../../constant/serviceURL';

export const forgetPasswordService = async (email: string) => {
  try {
    const response = await axios.post(`${authURL}/forget-password/${email}`);
    return response.data.status;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status === 400) {
      returnError = {
        message: ERR_Messages.NOT_FOUND_TOKEN,
        status: 400,
        toastStatus: ToastStatus.ERROR,
      };
    } else if (requestError.status === 404) {
      returnError = {
        message: ERR_Messages.NOT_FOUND,
        status: 404,
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


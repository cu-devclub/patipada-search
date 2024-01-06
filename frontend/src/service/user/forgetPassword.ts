import axios from 'axios';
import { CreateCustomError, ReturnError } from '../error';
import { ToastStatus, ERR_Messages } from '../../constant';

export const forgetPassword = async (email: string) => {
  try {
    //TODO : Test the environment mode
    const apiUrl =
      import.meta.env.MODE === "production"
        ? "http://auth-service:8082"
        : "http://localhost:8082";
    const response = await axios.post(`${apiUrl}/forget-password/${email}`);
    return response.data.status;
  } catch (error: unknown) {
     const requestError = CreateCustomError(error);
     let returnError: ReturnError;
     if (requestError.status === 400) {
       returnError = {
         message: ERR_Messages.NOT_FOUND_TOKEN,
         status: 404,
         toastStatus: ToastStatus.ERROR,
       };
     } else if (requestError.status === 404) {
       returnError = {
         message: ERR_Messages.NOT_FOUND_TOKEN,
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


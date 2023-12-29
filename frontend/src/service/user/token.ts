import axios from "axios";
import { CreateCustomError, ReturnError } from "../error";
import { ERR_Messages, ToastStatus } from "../../constant";
// use in contributor to verify their token
export const verifyToken = async (token: string) => {
  try {
    const apiUrl =
      import.meta.env.MODE === "production"
        ? import.meta.env.VITE_AUTH_API_URL
        : "http://localhost:8082";
    const response = await axios.get(`${apiUrl}/verify-token`, {
      headers: {
        Authorization: token,
      },
    });
    return response.data.result;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status == 400) {
      returnError = {
        message: ERR_Messages.MISSING_TOKEN,
        status: 400,
        toastStatus: ToastStatus.WARNING,
      };
    } else if (requestError.status === 401) {
      returnError = {
        message: ERR_Messages.INVALID_TOKEN,
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

// use in admin task to verify their token and role
export const authorize = async (token: string, requireRole: string) => {
  try {
    const apiUrl =
      import.meta.env.MODE === "production"
        ? import.meta.env.VITE_AUTH_API_URL
        : "http://localhost:8082";
    const response = await axios.get(
      `${apiUrl}/authorize?requiredRole=${requireRole}`,
      {
        headers: {
          Authorization: token,
        },
      }
    );

    return response.data.result;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status == 400) {
      returnError = {
        message: ERR_Messages.MISSING_TOKEN,
        status: 400,
        toastStatus: ToastStatus.WARNING,
      };
    } else if (requestError.status === 401) {
      returnError = {
        message: ERR_Messages.INVALID_TOKEN,
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

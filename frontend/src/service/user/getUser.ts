import axios from "../axiosInstance";
import { CreateCustomError, ReturnError } from "../error";
import { ToastStatus, ERR_Messages } from "../../constant";
import { authURL } from "../../constant/serviceURL";
import { User } from "../../models/user";
import { getCookie } from "typescript-cookie";

export const getAllUsers = async (): Promise<User[]> => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie("token");

    const response = await axios.get(`${authURL}/users`);
    if (response.data.users == null) {
      const users: User[] = [];
      return users;
    }

    const users = response.data.users.map((user: User) => {
      return {
        id: user.id,
        username: user.username,
        email: user.email,
        role: user.role,
      } as User;
    });

    return users;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status === 401) {
      returnError = {
        message: ERR_Messages.INVALID_TOKEN,
        status: 401,
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

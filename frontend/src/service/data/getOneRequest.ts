import axios from "axios";
import { ERR_Messages, ToastStatus } from "../../constant";
import { CreateCustomError, ReturnError } from "../error";
import { mapResponseToRequest,Request } from "../../models/request";
export const getRequestByRecordIndex = async (id: string): Promise<Request> => {
  try {
    const apiUrl =
      import.meta.env.MODE === "production"
        ? "http://data-service:8083"
        : "http://localhost:8083";
    const response = await axios.get(`${apiUrl}/requests/record/${id}`);
    const res = mapResponseToRequest(response.data.data);
    return res;
  } catch (error: unknown) {
    const requestError = CreateCustomError(error);
    let returnError: ReturnError;
    if (requestError.status === 404) {
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

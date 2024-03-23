import axios from "../axiosInstance";
import { getCookie } from "typescript-cookie";
import { CreateCustomError } from "../error";
import { Request } from "../../models/request";
import { dataURL } from "../../constant/serviceURL";

export const updateRequestService = async (data: Request) => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie("token");
    const response = await axios.put(`${dataURL}/request`, data);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

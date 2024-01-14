import axios from '../axiosInstance';
import { ERR_Messages, ToastStatus } from "../../constant";
import { CreateCustomError, ReturnError } from "../error";
import { mapResponseToRequest,Request } from "../../models/request";
import { getCookie } from 'typescript-cookie';
import { dataURL } from '../../constant/serviceURL';
export const getRequestByRecordIndex = async (id: string): Promise<Request> => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie('token');
    const response = await axios.get(`${dataURL}/request/latestRecord?index=${id}`);
    
    if (response.data.data == null) {
      const emptyRequest : Request =  {
        id: "",
        requestID: "NOT FOUND",
        index: "",
        question: "",
        answer: "",
        startTime: "",
        endTime: "",
        youtubeURL: "",
        status: "",
        createdAt: "",
        updatedAt: "",
        by: "",
        ApproveBy: "",
      };
      return emptyRequest;
    }

    const res = mapResponseToRequest(response.data.data);
    return res;
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

export const getRequestByParams = async(params: {status?: string, username?: string, requestID?: string, index?: string, approvedBy?: string}): Promise<Request[]> => {
  try {

    axios.defaults.headers.common["Authorization"] = getCookie('token');        

    const query = Object.keys(params)
      .filter((key: string) => params[key as keyof typeof params] !== undefined)
      .map((key: string) => `${key}=${encodeURIComponent(params[key as keyof typeof params] as string)}`)
      .join('&');

    const response = await axios.get(`${dataURL}/requests?${query}`);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const res = response.data.data.map((item: any) => mapResponseToRequest(item));
    return res;
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
}

///


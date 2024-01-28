import axios from '../axiosInstance';
import { getCookie } from 'typescript-cookie';
import { CreateCustomError } from "../error";
import { InsertRequestModels } from "../../models/request";
import { dataURL } from '../../constant/serviceURL';

export const insertRequest = async (data: InsertRequestModels) => {
  try {
    axios.defaults.headers.common["Authorization"] = getCookie('token');        
    const response = await axios.post(`${dataURL}/requests`, data);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

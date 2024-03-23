import axios from "../axiosInstance";
import { CreateCustomError } from "../error";
import { Rating } from "../../models/ratings";
import { dataURL } from "../../constant/serviceURL";

export const insertRatings = async (data: Rating) => {
  try {
    const response = await axios.post(`${dataURL}/ratings`, data);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

export const getAverageRatings = async () => {
  try {
    const response = await axios.get(`${dataURL}/ratings/average`);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

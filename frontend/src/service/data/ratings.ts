import axios from "axios";
import { CreateCustomError } from "../error";
import { Rating } from "../../models/ratings";
import { dataURL } from "../../constant/serviceURL";

export const insertRatingsService = async (data: Rating) => {
  try {
    const response = await axios.post(`${dataURL}/ratings`, data);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

export const getSummaryRatingsService = async () => {
  try {
    const response = await axios.get(`${dataURL}/ratings/average`);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

export const getRatingsService = async () => {
  try {
    const response = await axios.get(`${dataURL}/ratings`);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

import axios from "axios";

import { CreateCustomError } from "../error";
import { InsertRequestModels } from "../../models/request";

export const insertRequest = async (data: InsertRequestModels) => {
  try {
    // TODO : Attach token
    const apiUrl =
      import.meta.env.MODE === "production"
        ? "http://data-service:8083"
        : "http://localhost:8083";
    const response = await axios.post(`${apiUrl}/requests`, data);
    return response.data;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

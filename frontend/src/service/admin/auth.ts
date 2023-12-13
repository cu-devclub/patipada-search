import axios from "axios";
import { CustomError } from "../error";
export const login = async (username: string, password: string) => {
    try {
        //TODO : Test the environment mode 
        const apiUrl = import.meta.env.MODE === 'production' ? import.meta.env.VITE_AUTH_API_URL : "http://localhost:8082";
        const response = await axios.post(`${apiUrl}/login`, {
            username,
            password,
        });
        return response.data;
   } catch (error: unknown) {
        if (axios.isAxiosError(error)) {
            const customError: CustomError = {
                    message: `Axios error: ${error.message?.toString()}`,
                    status: error.response?.status,
            };
            throw customError;
        } else {
            const customError: CustomError = {
                message: `Unknown error: ${error}`,
            };
            throw customError;
        }
    }
}
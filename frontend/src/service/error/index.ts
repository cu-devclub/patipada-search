import axios from "axios";
import { ToastStatusType } from "../../constant";
interface CustomError {
    message: string;
    status?: number;
}

export interface ReturnError {
  message: string;
  status: number;
  toastStatus: ToastStatusType;
}

/**
 * Creates a custom error based on the given input error.
 *
 * @param {unknown} error - The error object to create a custom error from.
 * @return {CustomError} The created custom error.
 */

export const CreateCustomError = (error : unknown): CustomError => {
    if (axios.isAxiosError(error)) {
        const customError: CustomError = {
                message: `${error.response?.data.message}`,
                status: error.response?.status,
        };
        return customError;
    } else {
        const customError: CustomError = {
            message: `Unknown error: ${error}`,
            status: 500,
        };
        return customError;
    }
}
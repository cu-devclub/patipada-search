import axios from "axios";

interface CustomError {
    message: string;
    status?: number;
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
                message: `Axios error: ${error.message?.toString()}`,
                status: error.response?.status,
        };
        return customError;
    } else {
        const customError: CustomError = {
            message: `Unknown error: ${error}`,
        };
        return customError;
    }
}
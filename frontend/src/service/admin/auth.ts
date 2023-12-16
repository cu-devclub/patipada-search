import axios from "axios";
import { CreateCustomError } from "../error";

/**
 * Authenticates a user by sending a login request to the server.
 *
 * @param {string} username - The username of the user.
 * @param {string} password - The password of the user.
 * @return {Promise<any>} - A promise that resolves with the response data from the server.
 */
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
        const returnError = CreateCustomError(error);
        throw returnError;
    }
}
import axios from "axios";
import { CreateCustomError } from "../error";

/**
 * Performs a search query using the specified query string.
 *
 * @param {string} query - The search query string.
 * @return {Promise<any>} - A promise that resolves with the search results.
 */
export const search = async (query: string) => {
    try {
        //TODO : Test the environment mode 
        const apiUrl = import.meta.env.MODE === 'production' ? import.meta.env.VITE_SEARCH_API_URL : "http://localhost:8081";
        const response = await axios.get(`${apiUrl}/search?query=${query}`);
        return response.data;
   } catch (error: unknown) {
       const returnErr = CreateCustomError(error);
       throw returnErr;
    }
}
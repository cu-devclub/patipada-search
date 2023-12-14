import axios from "axios";
import { CreateCustomError } from "../error";
import { DataItem,SearchResultInterface} from "../../models/qa";
/**
 * Performs a search query using the specified query string.
 *
 * @param {string} query - The search query string.
 * @return {Promise<SearchResultInterface>} - A promise that resolves with the search results.
 */
export const search = async (query: string) => {
    try {
        //TODO : Test the environment mode 
        const apiUrl = import.meta.env.MODE === 'production' ? import.meta.env.VITE_SEARCH_API_URL : "http://localhost:8081";
        const response = await axios.get(`${apiUrl}/search?query=${query}`);
        
        const records: DataItem[] = response.data.results.map((item: DataItem) => {
            return {
                index: item.index,
                youtubeURL: item.youtubeURL,
                question: item.question,
                answer: item.answer,
                startTime: item.startTime,
                endTime: item.endTime,
            } as DataItem;
        });

        const searchResult: SearchResultInterface = {
            data: records,
            query: query,
            tokens: response.data.tokens,
        };

        return searchResult;
   } catch (error: unknown) {
       const returnErr = CreateCustomError(error);
       throw returnErr;
    }
}
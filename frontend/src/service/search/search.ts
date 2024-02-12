import axios from "../axiosInstance";
import { CreateCustomError } from "../error";
import { DataItem, SearchResultInterface } from "../../models/qa";
import { searchURL } from "../../constant/serviceURL";
/**
 * Performs a search query using the specified query string.
 *
 * @param {string} query - The search query string.
 * @param {string} searchType - Optional. The search type to use. Defaults to "tf_idf".
 * @return {Promise<SearchResultInterface>} - A promise that resolves with the search results.
 */
export const search = async (
  query: string,
  searchType: string = "tf_idf"
): Promise<SearchResultInterface> => {
  try {
    const response = await axios.get(
      `${searchURL}/search?query=${query}&searchType=${searchType}`
    );

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
};

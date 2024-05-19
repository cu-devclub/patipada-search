import axios from "../axiosInstance";
import { CreateCustomError } from "../error";
import { DataItem, SearchResultInterface } from "../../models/qa";
import { searchURL } from "../../constant/serviceURL";
import {
  SEARCH_STATUS,
  SEARCH_TYPE,
  SearchResultItemsPerPage,
  SearchResultPageNumbers,
} from "../../constant";
/**
 * Performs a search query using the specified query string.
 *
 * @param {string} query - The search query string.
 * @param {string} searchType - Optional. The search type to use. Defaults to "tf_idf".
 * @param {string} searchStatus - Optional. The search status to use. Defaults to "draft".
 * @param {number} offset - Optional. The offset of the search results. Defaults to 0.
 * @param {number} amount - Optional. The limit of the search results. Defaults to 8.
 * @return {Promise<SearchResultInterface>} - A promise that resolves with the search results.
 */
export const searchService = async (
  query: string,
  searchType: string = SEARCH_TYPE.DEFAULT,
  searchStatus: string = SEARCH_STATUS.DEFAULT,
  offset: number = 0,
  amount: number = SearchResultItemsPerPage,
  countNeeded: boolean = true,
  numPages: number = 0
): Promise<SearchResultInterface> => {
  try {
    const response = await axios.get(
      `${searchURL}/search?query=${query}&searchType=${searchType}&searchStatus=${searchStatus}&offset=${offset}&amount=${amount}&countNeeded=${countNeeded}`
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

    if (countNeeded) {
      // count Need means new search so we have to update the number of pages
      numPages = Math.ceil(response.data.amount / amount);
      if (numPages > SearchResultPageNumbers)
        numPages = SearchResultPageNumbers;
    }

    const searchResult: SearchResultInterface = {
      data: records,
      query: query,
      tokens: response.data.tokens,
      numPages: numPages,
    };

    return searchResult;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

export const searchRecordIndexService = async (
  recordIndex: string
): Promise<DataItem> => {
  try {
    const response = await axios.get(`${searchURL}/search/${recordIndex}`);
    const item: DataItem = {
      index: response.data.index,
      youtubeURL: response.data.youtubeURL,
      question: response.data.question,
      answer: response.data.answer,
      startTime: response.data.startTime,
      endTime: response.data.endTime,
    };
    return item;
  } catch (error: unknown) {
    const returnErr = CreateCustomError(error);
    throw returnErr;
  }
};

import { isValidEmail, isValueExist, isLengthEnough } from "./form";
import { timeToSeconds, splitTime, generateTime, formatDate } from "./time";
import { handleEnterKeyPress } from "./keyboard";
import { SignOut, ValidateToken, AuthorizeAdmin } from "./user";
import {
  extractStringFromHTML,
  encodeHTMLText,
  decodeHTMLText,
  convertStatusWord,
  checkIfCommentLeft,
} from "./request";
import {
  getStartAndEndIndexOfComments,
  removeCommentFromHTML,
} from "./comment";
import { twoDecimal } from "./rounding";
export {
  isValidEmail,
  isValueExist,
  isLengthEnough,
  timeToSeconds,
  handleEnterKeyPress,
  SignOut,
  ValidateToken,
  AuthorizeAdmin,
  extractStringFromHTML,
  splitTime,
  generateTime,
  encodeHTMLText,
  decodeHTMLText,
  convertStatusWord,
  checkIfCommentLeft,
  getStartAndEndIndexOfComments,
  removeCommentFromHTML,
  formatDate,
  twoDecimal,
};
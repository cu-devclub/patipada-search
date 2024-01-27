import { isValidEmail, isValueExist, isLengthEnough } from "./form";
import { timeToSeconds, splitTime, generateTime } from "./time";
import { handleEnterKeyPress } from "./keyboard";
import { SignOut, ValidateToken, AuthorizeAdmin } from "./user";
import {
  extractStringFromHTML,
  encodeHTMLText,
  decodeHTMLText,
  convertStatusWord,
  checkIfCommentLeft,
} from "./request";
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
};

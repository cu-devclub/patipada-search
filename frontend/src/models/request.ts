import { encodeHTMLText, decodeHTMLText } from "../functions";
import { DataItem } from "./qa";
export interface Request {
  id: string;
  requestID: string;
  index: string;
  question: string;
  answer: string;
  startTime: string;
  endTime: string;
  youtubeURL: string;
  status: string;
  createdAt: string;
  updatedAt: string;
  by: string;
  approvedBy: string;
}

export interface InsertRequestModels {
  index: string;
  question: string;
  answer: string;
  startTime: string;
  endTime: string;
  youtubeURL: string;
  by: string;
}

export interface RequestSummary {
  requestAmount: number;
  pendingAmount: number;
  reviewedAmount: number;
}

export const mapRequestToInsertRequestModels = (
  data: Request
): InsertRequestModels => {
  return {
    index: data.index,
    question: encodeHTMLText(data.question),
    answer: encodeHTMLText(data.answer),
    startTime: encodeHTMLText(data.startTime),
    endTime: encodeHTMLText(data.endTime),
    youtubeURL: data.youtubeURL,
    by: data.by,
  };
};

export const createEncodeRequest = (data: Request): Request => {
  return {
    ...data,
    question: encodeHTMLText(data.question),
    answer: encodeHTMLText(data.answer),
    startTime: encodeHTMLText(data.startTime),
    endTime: encodeHTMLText(data.endTime),
  };
};

export const mapDataItemToRequest = (data: DataItem): Request => {
  return {
    id: "",
    requestID: "",
    index: data.index,
    question: data.question,
    answer: data.answer,
    startTime: data.startTime,
    endTime: data.endTime,
    youtubeURL: data.youtubeURL,
    status: "",
    createdAt: "",
    updatedAt: "",
    by: "",
    approvedBy: "",
  };
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const mapResponseToRequest = (data: any): Request => {
  return {
    id: data.id,
    requestID: data.requestID,
    index: data.index,
    question: decodeHTMLText(data.question),
    answer: decodeHTMLText(data.answer),
    startTime: decodeHTMLText(data.startTime),
    endTime: decodeHTMLText(data.endTime),
    youtubeURL: data.youtubeURL,
    status: data.status,
    createdAt: data.createdAt,
    updatedAt: data.updatedAt,
    by: data.by,
    approvedBy: data.approvedBy,
  };
};

export const mockData: Request = {
  id: "1",
  requestID: "1",
  index: "1",
  question: `<p>ผมตั้งใจกําหนดรู้ชัดในขันธ์ 5 ให้ติดต่อเนื่อง รู้ว่าในแต่ละขันธ์ แต่ละคขันธ์ทํางานแยกกันไม่มีเราอยู่ในแต่ละขันธ์ทั้ง 5 ผมพิจารณาต่อเนื่องในระหว่างนั่งอานาแล้วแล้วจู่ ๆ ผมสงสัยและถามเข้าไปในจิตว่าแล้วการพิจารณาอริยสัจ 4 เห็นเกิดดับละได้พิจารณาไหม เพราะผมตั้งใจกําหนดรู้ในขันธ์ 5 อย่างเดียวเลยครับ พอสิ้นคําถามอยู่ ๆ ก็เริ่มสะอื้นน้ำตาคลอเพราะ<span data-comment="{&quot;uuid&quot;:&quot;221e29ef-96e5-4cee-85f9-f0dcef73af41&quot;,&quot;comments&quot;:[{&quot;userName&quot;:&quot;super-admin&quot;,&quot;time&quot;:1704555760459,&quot;content&quot;:&quot;sdfsdf&quot;}]}">จิตต</span>อบกลับมาเองว่าก็ที่รู้ชัดในขันธ์ 5 ได้ขณะใดคือขณะที่รู้อริยสัจเห็นเกิดดับอยู่แล้วความเห็นนี้ถูกต้องไหมครับพระอาจารย์ ขอคําแนะนําการปฏิบัติจากพระอาจารย์ด้วยขอรับ เพราะตอนนี้ผมตั้งใจปฏิบัติเพียรกําหนดรู้ชัดในขันธ์ 5 ให้ติดต่อเนื่องในชีวิตประจําวันให้ได้ทุกขณะครับ</p>`,
  answer:
    "TypeScript is a typed superset of JavaScript that compiles to plain JavaScript.",
  startTime: "00:00:10",
  endTime: "00:00:30",
  youtubeURL: "61oREuQ5JU8",
  status: "pending",
  createdAt: "2021-06-01T00:00:00.000Z",
  updatedAt: "2021-06-01T00:00:00.000Z",
  by: "1",
  approvedBy: "1",
};

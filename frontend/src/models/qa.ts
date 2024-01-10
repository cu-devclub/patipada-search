export interface DataItem {
  index: string;
  question: string;
  answer: string;
  startTime: string;
  endTime: string;
  youtubeURL: string;
}

export const mockData: DataItem = {
  index: "1",
  question:
    "What is TypeScript? What is TypeScript? What is TypeScript? What is TypeScript? What is TypeScript? What is TypeScript?",
  answer:
    "TypeScript is a typed superset of JavaScript that compiles to plain JavaScript.",
  startTime: "00:00:10",
  endTime: "00:00:30",
  youtubeURL: "61oREuQ5JU8",
};

export interface SearchResultInterface {
  data: DataItem[];
  query: string;
  tokens: string[];
}

export const mapResponseToDataItem = (data): DataItem => {
  return {
    index: data.index,
    question: data.question,
    answer: data.answer,
    startTime: data.startTime,
    endTime: data.endTime,
    youtubeURL: data.youtubeURL,
  };
};

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
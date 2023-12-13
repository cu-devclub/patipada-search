export interface DataItem {
    id: number;
    question: string;
    answer: string;
    startTime: string;
    endTime: string;
    youtubeURL: string;
}

export interface SearchResultInterface {
    data: DataItem[];
    query: string;
    tokens: string[];
}
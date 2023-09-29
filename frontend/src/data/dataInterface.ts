export interface DataItem {
    id: number;
    question: string;
    answer: string;
    startTime: string;
    endTime: string;
    youtubeURL: string;
}

export interface DataProps {
    data: DataItem;
}

export interface QAProps {
    data: DataItem;
    query: string;
}

export interface SearchResultsProps {
    data: DataItem[]; // Ensure that the type matches your data structure
    query: string;
}

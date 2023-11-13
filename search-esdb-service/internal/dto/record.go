package dto

import "fmt"

type QARecord struct {
	Index      string    `json:"index"`
	YoutubeURL string `json:"youtubeURL"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

func (q *QARecord) ToString() string {
	return fmt.Sprintf("URL %v Question %v Answer %v StartTime %v EndTime %v", q.YoutubeURL, q.Question, q.Answer, q.StartTime, q.EndTime)
}

type Token struct {
	Token       string `json:"token"`
	StartOffset int    `json:"start_offset"`
	EndOffset   int    `json:"end_offset"`
	Type        string `json:"type"`
	Position    int    `json:"position"`
}

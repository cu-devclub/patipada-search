package dto

import "fmt"

type QARecord struct {
	YoutubeURL string `json:"youtubeURL"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

func (q *QARecord) ToString() string {
	return fmt.Sprintf("URL %v Question %v Answer %v StartTime %v EndTime %v", q.YoutubeURL, q.Question, q.Answer, q.StartTime, q.EndTime)
}

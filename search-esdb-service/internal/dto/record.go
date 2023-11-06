package dto

import "fmt"

type QARecord struct {
	YoutubeURL string
	Question   string
	Answer     string
	StartTime  string
	EndTime    string
}

func (q *QARecord) ToString() string {
	return fmt.Sprintf("URL %v Question %v Answer %v StartTime %v EndTime %v", q.YoutubeURL, q.Question, q.Answer, q.StartTime, q.EndTime)
}

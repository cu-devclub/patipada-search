// entity for record domain
// entity is a data model use to deals with data
// in the low level e.g. data in db or migration files

package entities

import "fmt"

type (
	Record struct {
		Index       string    `json:"index"`
		YoutubeURL  string    `json:"youtubeURL"`
		Question    string    `json:"question"`
		Answer      string    `json:"answer"`
		StartTime   string    `json:"startTime"`
		EndTime     string    `json:"endTime"`
		QuestionLDA []float64 `json:"question_lda,omitempty"`
		AnswerLDA   []float64 `json:"answer_lda,omitempty"`
	}

	UpdateRecord struct {
		DocumentID string `json:"documentID"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
		Question   string `json:"question"`
		Answer     string `json:"answer"`
	}

	Token struct {
		Token       string `json:"token"`
		StartOffset int    `json:"start_offset"`
		EndOffset   int    `json:"end_offset"`
		Type        string `json:"type"`
		Position    int    `json:"position"`
	}

	Text2VecResponse struct {
		Embedding   []float32 `json:"Embedding"`
		Name        string    `json:"Name"`
		ScoreWeight float32   `json:"ScoreWeight"`
	}
)

func (r *Record) ToString() string {
	return fmt.Sprintf("Index: %s, YoutubeURL: %s, Question: %s, Answer: %s, StartTime: %s, EndTime: %s, QuestionLDA: %v, AnswerLDA: %v", r.Index, r.YoutubeURL, r.Question, r.Answer, r.StartTime, r.EndTime, r.QuestionLDA, r.AnswerLDA)
}

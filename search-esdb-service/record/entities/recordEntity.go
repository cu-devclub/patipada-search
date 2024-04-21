// entity for record domain
// entity is a data model use to deals with data
// in the low level e.g. data in db or migration files

package entities

import (
	"encoding/json"
	"fmt"
)

type (
	Record struct {
		Index      string           `json:"index"`
		YoutubeURL string           `json:"youtubeURL"`
		Question   string           `json:"question"`
		Answer     string           `json:"answer"`
		StartTime  string           `json:"startTime"`
		EndTime    string           `json:"endTime"`
		Vectors    []*RecordVectors `json:"vectors"`
	}

	RecordVectors struct {
		Model   string   `json:"model"`
		Vectors *Vectors `json:"vectors"`
	}

	Vectors struct {
		RecordIndex    string    `json:"recordIndex"`
		QuestionVector []float32 `json:"question_vector"`
		AnswerVector   []float32 `json:"answer_vector"`
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
)

func (r *Record) ToString() string {
	return fmt.Sprintf("Index: %s, YoutubeURL: %s, Question: %s, Answer: %s, StartTime: %s, EndTime: %s", r.Index, r.YoutubeURL, r.Question, r.Answer, r.StartTime, r.EndTime)
}

func (r *Record) BuildJson() ([]byte, error) {
	// Create a map to hold the JSON fields
	jsonFields := make(map[string]interface{})

	// Add the static fields to the map
	jsonFields["index"] = r.Index
	jsonFields["youtubeURL"] = r.YoutubeURL
	jsonFields["question"] = r.Question
	jsonFields["answer"] = r.Answer
	jsonFields["startTime"] = r.StartTime
	jsonFields["endTime"] = r.EndTime

	// Add the dynamic fields to the map
	for _, vector := range r.Vectors {
		model := vector.Model
		jsonFields[model+"-question"] = vector.Vectors.QuestionVector
		jsonFields[model+"-answer"] = vector.Vectors.AnswerVector
	}

	// Marshal the map to JSON
	jsonData, err := json.Marshal(jsonFields)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

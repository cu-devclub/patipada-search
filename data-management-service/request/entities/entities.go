package entities

import (
	"data-management/util"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Request struct {
		ID         string    `bson:"_id,omitempty" json:"id,omitempty"`
		RequestID  string    `bson:"request_id"`
		Index      string    `bson:"index"`
		YoutubeURL string    `bson:"youtubeURL"`
		Question   string    `bson:"question"`
		Answer     string    `bson:"answer"`
		StartTime  string    `bson:"startTime"`
		EndTime    string    `bson:"endTime"`
		CreatedAt  time.Time `bson:"created_at"`
		UpdatedAt  time.Time `bson:"updated_at"`
		Status     string    `bson:"status"` // "pending", "reviewed"
		By         string    `bson:"by"`
		ApprovedBy string    `bson:"approved_by"`
	}

	Filter struct {
		Status     string `bson:"status"`
		By         string `bson:"by"`
		RequestID  string `bson:"request_id"`
		Index      string `bson:"index"`
		ApprovedBy string `bson:"approved_by"`
	}

	Record struct {
		Index      string `json:"index"`
		YoutubeURL string `json:"youtubeURL"`
		Question   string `json:"question"`
		Answer     string `json:"answer"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
	}
)

func (r *Request) MockData() {
	r.Index = "asdads-3"
	r.YoutubeURL = "https://www.youtube.com/watch?v=JGwWNGJdvx8"
	r.Question = "What is the name of the main character?"
	r.Answer = "Harry Potter"
	r.StartTime = "00:00:00"
	r.EndTime = "00:00:10"
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	r.Status = "pending"
	r.By = "user1"
}

func (f *Filter) ConvertToBsonM() (bson.M, error) {
	data, err := bson.Marshal(f)
	if err != nil {
		return nil, err
	}

	var m bson.M
	err = bson.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	// Remove empty fields from filter
	for key, value := range m {
		if value == "" {
			delete(m, key)
		}
	}

	return m, nil
}

func (r *Record) ExtractHTML() {
	r.StartTime = util.ExtractRawStringFromHTMLTags(r.StartTime)
	r.EndTime = util.ExtractRawStringFromHTMLTags(r.EndTime)
	r.Question = util.ExtractRawStringFromHTMLTags(r.Question)
	r.Answer = util.ExtractRawStringFromHTMLTags(r.Answer)
}

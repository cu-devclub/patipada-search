package models

import (
	"time"
)

type (
	Request struct {
		ID         string    `json:"id,omitempty"`
		RequestID  string    `json:"request_id"`
		Index      string    `json:"index" binding:"required" validate:"required"`
		YoutubeURL string    `json:"youtubeURL" binding:"required" validate:"required"`
		Question   string    `json:"question" binding:"required" validate:"required"`
		Answer     string    `json:"answer" binding:"required" validate:"required"`
		StartTime  string    `json:"startTime" binding:"required" validate:"required"`
		EndTime    string    `json:"endTime" binding:"required" validate:"required"`
		CreatedAt  time.Time `json:"created_at,omitempty"`
		UpdatedAt  time.Time `json:"updated_at,omitempty"`
		Status     string    `json:"status" validate:"omitempty,oneof=pending reviewed"` // "pending", "reviewed"
		By         string    `json:"by" binding:"required" validate:"required"`
		ApprovedBy string    `json:"approved_by,omitempty"`
		CommentUID string    `json:"comment_uid,omitempty"`
	}

	Summary struct {
		RecordSummary  *RecordSummary  `json:"recordSummary"`
		RequestSummary *RequestSummary `json:"requestSummary"`
	}

	RecordSummary struct {
		RecordAmount      int `json:"recordAmount"`
		YouTubeClipAmount int `json:"youtubeClipAmount"`
	}

	RequestSummary struct {
		RequestAmount  int `json:"requestAmount"`
		ReviewedAmount int `json:"reviewedAmount"`
		PendingAmount  int `json:"pendingAmount"`
	}
)

func (r *Request) ToString() string {
	return r.Index + " " + r.YoutubeURL + " " + r.Question + " " + r.Answer + " " + r.StartTime + " " + r.EndTime + " " + r.By
}

func (r *Request) MockData() {
	r.Index = "61oREuQ5JU8-1"
	r.YoutubeURL = "https://www.youtube.com/watch?v=JGwWNGJdvx8"
	r.Question = "What is the name of the main character?"
	r.Answer = "Harry Potter"
	r.StartTime = "00:00:00"
	r.EndTime = "00:00:10"
	r.By = "admin"
}

func (r *Request) CreateMockJSON() string {
	s := `{
		"index": "asdads-3",
		"youtubeURL": "https://www.youtube.com/watch?v=JGwWNGJdvx8",
		"question": "What is the name of the main character?",
		"answer": "Harry Potter",
		"startTime": "00:00:00",
		"endTime": "00:00:10",
		"by": "admin"
	}`
	return s
}

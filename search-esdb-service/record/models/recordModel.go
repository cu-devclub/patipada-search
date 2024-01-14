// models for record domain
// model is a data model use to deals with data
// in the high level e.g. response to frontend

package models

type (
	Record struct {
		Index      string `json:"index"`
		YoutubeURL string `json:"youtubeURL"`
		Question   string `json:"question"`
		Answer     string `json:"answer"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
	}

	UpdateRecord struct {
		DocumentID string `json:"documentID"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
		Question   string `json:"question"`
		Answer     string `json:"answer"`
	}

	SearchRecordStruct struct {
		Results []*Record `json:"results"`
		Tokens  []string  `json:"tokens"`
	}
)

func (s *SearchRecordStruct) ToString() string {
	return `{
		"results": ` + s.ResultsToString() + `,
		"tokens": ` + s.TokensToString() + `
	}`
}

func (s *SearchRecordStruct) ResultsToString() string {
	var results string
	for _, result := range s.Results {
		results += result.ToString() + ","
	}
	return "[" + results[:len(results)-1] + "]"
}

func (r *Record) ToString() string {
	return `{
		"index": "` + r.Index + `",
		"youtubeURL": "` + r.YoutubeURL + `",
		"question": "` + r.Question + `",
		"answer": "` + r.Answer + `",
		"startTime": "` + r.StartTime + `",
		"endTime": "` + r.EndTime + `"
	}`
}

func (s *SearchRecordStruct) TokensToString() string {
	var tokens string
	for _, token := range s.Tokens {
		tokens += `"` + token + `",`
	}
	return "[" + tokens[:len(tokens)-1] + "]"
}

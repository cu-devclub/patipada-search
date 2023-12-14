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

	SearchRecordStruct struct {
		Results []*Record `json:"results"`
		Tokens []string  `json:"tokens"`
	}
)

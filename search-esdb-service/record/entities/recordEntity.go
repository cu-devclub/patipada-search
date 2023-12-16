// entity for record domain 
// entity is a data model use to deals with data
// in the low level e.g. data in db or migration files 

package entities

type (
	Record struct {
		Index      string `json:"index"`
		YoutubeURL string `json:"youtubeURL"`
		Question   string `json:"question"`
		Answer     string `json:"answer"`
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime"`
	}

	Token struct {
		Token       string `json:"token"`
		StartOffset int    `json:"start_offset"`
		EndOffset   int    `json:"end_offset"`
		Type        string `json:"type"`
		Position    int    `json:"position"`
	}
)

package migration

import (
	"encoding/json"
	"search-esdb-service/config"
)

type IndexCreationBody struct {
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}

type Settings struct {
	Index Index `json:"index"`
}

type Index struct {
	Analysis Analysis `json:"analysis"`
}

type Analysis struct {
	Analyzer map[string]Analyzer `json:"analyzer"`
	Filter   map[string]Filter   `json:"filter"`
}

type Analyzer struct {
	Tokenizer string   `json:"tokenizer"`
	Filter    []string `json:"filter"`
}

type Filter struct {
	Type           string `json:"type"`
	MaxShingleSize int    `json:"max_shingle_size"`
	MinShingleSize int    `json:"min_shingle_size"`
	OutputUnigrams bool   `json:"output_unigrams"`
}

type Mappings struct {
	Properties map[string]Property `json:"properties"`
}

type Property struct {
	Type     string `json:"type"`
	Analyzer string `json:"analyzer,omitempty"`
	Dims     int    `json:"dims,omitempty"`
}

func CreateIndexCreationBody(mlConfig *config.MLConfig) (string, error) {
	fields := createNeccessaryFields(mlConfig)
	indexCreationBody := IndexCreationBody{
		Settings: Settings{
			Index: Index{
				Analysis: Analysis{
					Analyzer: map[string]Analyzer{
						"analyzer_shingle": {
							Tokenizer: "icu_tokenizer",
							Filter:    []string{"filter_shingle"},
						},
					},
					Filter: map[string]Filter{
						"filter_shingle": {
							Type:           "shingle",
							MaxShingleSize: 3,
							MinShingleSize: 2,
							OutputUnigrams: true,
						},
					},
				},
			},
		},
		Mappings: Mappings{
			Properties: fields,
		},
	}

	jsonData, err := json.Marshal(indexCreationBody)
	if err != nil {
		return "", err
	}

	return string(jsonData),nil
}

func createNeccessaryFields(mlConfig *config.MLConfig) map[string]Property {
	fields := map[string]Property{
		"youtubeURL": {
			Type: "text",
		},
		"question": {
			Type:     "text",
			Analyzer: "analyzer_shingle",
		},
		"answer": {
			Type:     "text",
			Analyzer: "analyzer_shingle",
		},
		"startTime": {
			Type: "text",
		},
		"endTime": {
			Type: "text",
		},
	}

	for _, api := range mlConfig.APIs {
		fields[api.Name+"-question"] = Property{
			Type: "dense_vector",
			Dims: 30,
		}

		fields[api.Name+"-answer"] = Property{
			Type: "dense_vector",
			Dims: 30,
		}
	}

	return fields
}

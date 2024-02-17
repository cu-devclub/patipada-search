package repositories

import (
	"context"
	"fmt"
	"search-esdb-service/errors"
	"search-esdb-service/record/entities"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (r *RecordESRepository) UpdateRecord(record *entities.UpdateRecord) *errors.RequestError {
	// The partial document to update.
	doc := fmt.Sprintf(`{
		"doc": {
			"startTime": "%s",
			"endTime": "%s",
			"question": "%s",
			"answer": "%s"
		}
	}`, record.StartTime, record.EndTime, record.Question, record.Answer)

	// Initialize an Update API request object.
	req := esapi.UpdateRequest{
		Index:      "record",
		DocumentID: record.DocumentID,
		Body:       strings.NewReader(doc),
		Refresh:    "true",
	}

	// Execute the request.
	res, err := req.Do(context.Background(), r.es)
	if err != nil {
		return errors.CreateError(500, fmt.Sprintf("Error getting record: %s", err))
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.CreateError(res.StatusCode, fmt.Sprintf("Elasticsearch error: %s", res.Status()))
	}

	return nil

}

package repositories

import (
	"context"
	"fmt"
	"log"
	"search-esdb-service/record/entities"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (r *RecordESRepository) UpdateRecord(record *entities.UpdateRecord) error {
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
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return err
	}
	
	log.Println("Update record success...")
	return nil

}

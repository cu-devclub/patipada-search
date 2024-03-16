package repositories

import (
	"context"
	"data-management/constant"
	"data-management/request/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// UpdateRequest updates a request in the MongoDB collection.
//
// The function takes a pointer to an entities.Request object as input. The ID field of the Request object
// is used to find the document to update in the MongoDB collection. The other fields of the Request object
// are used to update the corresponding fields in the document.
//
// The function returns an error if the update operation fails. If the update operation is successful,
// the function returns nil.
//
// The function uses the UpdateOne method from the mongo package to perform the update operation. The UpdateOne
// method updates the first document that matches the filter in the MongoDB collection.
//
// The function uses a context with a timeout of 10 seconds to ensure that the update operation is cancelled
// if it takes too long.
//
// Example:
//
//	request := &entities.Request{
//	    ID: "60d5ecf7c88f9a200f9e2c5a",
//	    Question: "Updated question",
//	    Answer: "Updated answer"
//	    ....
//	}
//	err := r.UpdateRequest(request)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (r *requestRepositories) UpdateRequest(request *entities.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(request.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"question":    request.Question,
			"answer":      request.Answer,
			"startTime":   request.StartTime,
			"endTime":     request.EndTime,
			"status":      request.Status,
			"updated_at":  request.UpdatedAt,
			"approved_by": request.ApprovedBy,
		},
	}

	_, err = r.requestCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *requestRepositories) UpdateRecord(record *entities.Record) error {
	entity := &entities.UpdateRecord{
		DocumentID: record.Index,
		Question:   record.Question,
		Answer:     record.Answer,
		StartTime:  record.StartTime,
		EndTime:    record.EndTime,
	}
	err := r.communicationClient.PublishUpdateRecordsToRabbitMQ(
		constant.UPDATE_RECORD_PAYLOAD_NAME, 
		entity,
	)
	if err != nil {
		return err
	}
	return nil
}

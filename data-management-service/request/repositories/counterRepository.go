package repositories

import (
	"data-management/request/entities"

	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

// GetNextRequestCounter increments the sequence number in the "counters" collection in MongoDB.
// This function is used to generate a unique sequence number for each new Request.
// It returns the next sequence number and any error encountered.
//
// The function works by finding the document in the "counters" collection with _id "request",
// incrementing its "seq" field, and returning the updated "seq" value.
// If the document does not exist, MongoDB will return an error.
//
// Usage:
//
//	nextSeq, err := repositoryImpl.GetNextRequestCounter()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	nextRequestID := fmt.Sprintf("REQ%d", nextSeq)
//
// This will generate a unique RequestID for each new Request, like "REQ1", "REQ2", etc.
func (r *repositoryImpl) GetNextRequestCounter() (int, error) {
	var counter struct {
		Seq int `bson:"seq"`
	}
	err := r.requestCounterCollection.FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": "request"},
		bson.M{"$inc": bson.M{"seq": 1}},
		options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true),
	).Decode(&counter)
	if err != nil {
		return 0, err
	}
	return counter.Seq, nil
}

func (r *repositoryImpl) UpsertRecordCounter(recordCounter *entities.RecordCounter) error {
	_, err := r.requestCounterCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": "record"},
		bson.M{
			"$set": bson.M{
				"recordAmount":      recordCounter.RecordAmount,
				"youtubeClipAmount": recordCounter.YoutubeClipAmount,
			},
		},
		options.Update().SetUpsert(true),
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryImpl) GetRecordCounter() (*entities.RecordCounter, error) {
	var counter *entities.RecordCounter
	err := r.requestCounterCollection.FindOne(
		context.TODO(),
		bson.M{"_id": "record"},
	).Decode(&counter)
	if err != nil {
		return counter, err
	}
	return counter, nil
}

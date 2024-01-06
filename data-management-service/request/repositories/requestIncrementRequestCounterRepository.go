package repositories

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
)

// IncrementRequestCounter increments the sequence number in the "counters" collection in MongoDB.
// This function is used to generate a unique sequence number for each new Request.
// It returns the next sequence number and any error encountered.
//
// The function works by finding the document in the "counters" collection with _id "request",
// incrementing its "seq" field, and returning the updated "seq" value.
// If the document does not exist, MongoDB will return an error.
//
// Usage:
//
//	nextSeq, err := RequestRepositories.IncrementRequestCounter()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	nextRequestID := fmt.Sprintf("REQ%d", nextSeq)
//
// This will generate a unique RequestID for each new Request, like "REQ1", "REQ2", etc.
func (r *requestRepositories) IncrementRequestCounter() (int, error) {
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
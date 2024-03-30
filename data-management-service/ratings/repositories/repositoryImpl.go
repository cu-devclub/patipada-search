package repositories

import (
	"context"
	"data-management/errors"
	"data-management/ratings/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ratingsRepository struct {
	mongo             *mongo.Client
	ratingsCollection *mongo.Collection
}

func NewRatingRepository(mongo *mongo.Client) Repository {
	ratingsCollection := mongo.Database("data").Collection("ratings")
	return &ratingsRepository{
		mongo:             mongo,
		ratingsCollection: ratingsCollection,
	}
}

func (r *ratingsRepository) InsertRating(rating *entities.Rating) (string, error) {
	result, err := r.ratingsCollection.InsertOne(context.TODO(), rating)
	if err != nil {
		return "", err
	}

	// Get the inserted ID and assert it to an ObjectID
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.CreateError(500, "Cannot assert inserted ID to an ObjectID")
	}

	// Convert the ObjectID to a string
	objectIDString := objectID.Hex()

	return objectIDString, nil
}

func (r *ratingsRepository) GetRatings() ([]*entities.Rating, error) {
	cursor, err := r.ratingsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.CreateError(500, "Cannot get ratings")
	}
	defer cursor.Close(context.Background())

	ratings := make([]*entities.Rating, 0)
	for cursor.Next(context.Background()) {
		var rating entities.Rating
		err := cursor.Decode(&rating)
		if err != nil {
			return nil, errors.CreateError(500, "Cannot decode rating")
		}
		ratings = append(ratings, &rating)
	}

	return ratings, nil
}

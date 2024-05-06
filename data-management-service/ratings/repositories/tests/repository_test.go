package rating_repository_test

import (
	test_container_database "data-management/mock/testcontainer/database"
	"data-management/ratings/entities"
	"data-management/ratings/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRepoTest() (repositories.Repository,func()) {
	// setup
	db,cleanup, err := test_container_database.NewMockMongoClient()
	if err != nil {
		panic(err)
	}

	repo := repositories.NewRatingRepository(db)
	return repo,cleanup
}

func TestRatings(t *testing.T) {
	repo,cleanup := setupRepoTest()

	defer cleanup()

	t.Run("GetRatings no value", func(t *testing.T) {
		ratings, err := repo.GetRatings()
		assert.Nil(t, err)

		assert.Equal(t, 0, len(ratings))
	})
	
	t.Run("InsertRating", func(t *testing.T) {
		rating := &entities.Rating{
			Stars:   5,
			Comment: "Good",
		}
		res, err := repo.InsertRating(rating)
		assert.Nil(t, err)

		assert.NotEqual(t, "", res)
	})

	t.Run("GetRatings the value exist", func(t *testing.T) {
		expectedRating := &entities.Rating{
			RatingID: "1",
			Stars:    5,
			Comment:  "Good",
		}

		// insert the rating
		_, err := repo.InsertRating(expectedRating)
		assert.Nil(t, err)

		// get the rating
		ratings, err := repo.GetRatings()
		assert.Nil(t, err)

		// the ratings can have multiple we want to check that the rating we inserted is in the list
		found := false 
		for _, rating := range ratings {
			if rating.RatingID == expectedRating.RatingID {
				found = true
				assert.Equal(t, expectedRating.Stars, rating.Stars)
				assert.Equal(t, expectedRating.Comment, rating.Comment)
			}
		}

		assert.True(t, found)
	})


}
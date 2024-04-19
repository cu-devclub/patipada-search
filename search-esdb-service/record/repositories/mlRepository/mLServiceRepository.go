package repositories

import "search-esdb-service/communication"

type MLServiceRepository struct {
	comm communication.Communication
}

func NewMLServiceRepository(comm *communication.Communication) MLRepository {
	return &MLServiceRepository{
		comm: *comm,
	}
}

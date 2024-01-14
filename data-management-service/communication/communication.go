package communication

import "data-management/request/entities"

type Communication interface {
	Authorization(token string, requiredRole string) (bool, error)
	VerifyUsername(username string) (bool, error)
	SearchRecord(recordID string) (bool, error)
	UpdateRecord(record *entities.Record) (bool, error)
}

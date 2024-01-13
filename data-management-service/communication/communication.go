package communication

type Communication interface {
	Authorization(token string, requiredRole string) (bool, error)
	VerifyUsername(username string) (bool, error)
	SearchRecord(recordID string) (bool, error)
}

package communication

type Communication interface {

	// GRPC
	Authorization(token string, requiredRole string) (bool, error)
	VerifyUsername(username string) (bool, error)
	SearchRecord(recordID string) (bool, error)

	// RabbitMQ
	PublishUpdateRecordsToRabbitMQ(payloadName string, message interface{}) error
}

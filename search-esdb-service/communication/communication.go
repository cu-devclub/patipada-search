package communication

type Communication interface {
	// Rabbit MQ
	Listen(topics []string) error
}

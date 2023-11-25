package _interface

type ConsumerControllerInterface interface {
}

type ConsumerServiceInterface interface {
	Consume() error // Consume message from Kafka and Redis
}

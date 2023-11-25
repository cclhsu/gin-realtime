package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type ProducerControllerInterface interface {
	CreateTopic(ginContext *gin.Context) error // Create topic to Kafka and Redis
	DeleteTopic(ginContext *gin.Context) error // Delete topic to Kafka and Redis
	Produce(ginContext *gin.Context) error     // Produce message to Kafka and Redis
}

type ProducerServiceInterface interface {
	CreateTopic(topic string) error         // Create topic to Kafka and Redis
	DeleteTopic(topic string) error         // Delete topic to Kafka and Redis
	Produce(message model.MessageDTO) error // Produce message to Kafka and Redis
}

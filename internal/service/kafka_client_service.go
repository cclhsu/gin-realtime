package service

// go get -u github.com/confluentinc/confluent-kafka-go/kafka

import (
	"context"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaClientInterface interface {
	NewKafkaClientService(ctx context.Context, logger *logrus.Logger, bootstrapServers, groupID string, topics []string) (*KafkaClientService, error)
	_interface.ConsumerServiceInterface
	Initialize() error
	Shutdown() error
}

type KafkaClientService struct {
	ctx    context.Context
	logger *logrus.Logger
	// kafkaServerServiceURL string
	server   string
	GroupID  string
	consumer *kafka.Consumer
	topics   []string
}

func NewKafkaClientService(ctx context.Context, logger *logrus.Logger, bootstrapServers, groupID string, topics []string) (*KafkaClientService, error) {
	logger.Info("KafkaClientService NewKafkaClientService")
	logger.Debugf("KafkaClientService NewKafkaClientService: %v", bootstrapServers)
	logger.Debugf("KafkaClientService NewKafkaClientService: %v", topics)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	logger.Println("Connected to Kafka")

	return &KafkaClientService{
		ctx:      ctx,
		logger:   logger,
		server:   bootstrapServers,
		GroupID:  groupID,
		consumer: consumer,
		topics:   topics,
	}, nil
}

func (kcs *KafkaClientService) Initialize() error {
	kcs.logger.Info("KafkaClientService Initialize")
	go kcs.Consume()
	return nil
}

func (kcs *KafkaClientService) Shutdown() error {
	kcs.logger.Info("Shutting down KafkaClientService")
	if err := kcs.consumer.Close(); err != nil {
		kcs.logger.WithError(err).Error("Error shutting down Kafka consumer")
		return err
	}
	return nil
}

// func (kcs *KafkaClientService) initializeKafkaServerServiceURL() string {
// 	SERVER_HOST := os.Getenv("SERVER_HOST")
// 	if SERVER_HOST == "" {
// 		SERVER_HOST = "0.0.0.0"
// 	}
// 	SERVER_PORT := os.Getenv("SERVER_PORT")
// 	if SERVER_PORT == "" {
// 		SERVER_PORT = "3001"
// 	}
// 	return fmt.Sprintf("http://%s:%s/kafka/handler", SERVER_HOST, SERVER_PORT)
// }

func (kcs *KafkaClientService) Consume() error {
	kcs.logger.Info("KafkaClientService Consume: %v", kcs.topics)
	kcs.consumer.SubscribeTopics(kcs.topics, nil)

	for {
		select {
		case <-kcs.ctx.Done():
			kcs.logger.Info("KafkaClientService Consume stopped due to context cancellation")
			return nil
		default:
			msg, err := kcs.consumer.ReadMessage(-1)
			if err == nil {
				kcs.logger.Infof("Message on %s: %s", msg.TopicPartition, string(msg.Value))
			} else {
				kcs.logger.WithError(err).Errorf("Consumer error: %v (%v)", err, msg)
				return err
			}
		}
	}
}

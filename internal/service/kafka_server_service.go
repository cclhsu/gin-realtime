package service

// go get -u github.com/confluentinc/confluent-kafka-go/kafka

import (
	"context"
	"encoding/json"
	"fmt"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaServerInterface interface {
	NewKafkaServerService(ctx context.Context, logger *logrus.Logger, bootstrapServers string, topics []string) (*KafkaServerService, error)
	_interface.ProducerServiceInterface
	Initialize() error
}

type KafkaServerService struct {
	ctx      context.Context
	logger   *logrus.Logger
	producer *kafka.Producer
	topics   []string
}

func NewKafkaServerService(ctx context.Context, logger *logrus.Logger, bootstrapServers string, topics []string) (*KafkaServerService, error) {
	logger.Info("KafkaServerService NewKafkaServerService")
	logger.Debugf("KafkaServerService NewKafkaServerService: %v", bootstrapServers)
	logger.Debugf("KafkaServerService NewKafkaServerService: %v", topics)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}

	logger.Println("Connected to Kafka")

	return &KafkaServerService{
		ctx:      ctx,
		logger:   logger,
		producer: producer,
		topics:   topics,
	}, nil
}

func (kss *KafkaServerService) Initialize() error {
	kss.logger.Info("KafkaServerService Initialize")

	go func() {
		for e := range kss.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					kss.logger.WithError(ev.TopicPartition.Error).Error("Delivery failed")
				} else {
					kss.logger.Infof("Delivered message to %v", ev.TopicPartition)
				}
			}
		}
	}()

	return nil
}

func (kcs *KafkaServerService) CreateTopic(topic string) error {
	kcs.logger.Info("KafkaServerService CreateTopic")
	return fmt.Errorf("Not implemented")

	// err := kcs.producer.CreateTopics([]kafka.TopicSpecification{{
	// 	Topic:             topic,
	// 	NumPartitions:     1,
	// 	ReplicationFactor: 1,
	// }}, kafka.SetAdminOperationTimeout(1000))

	// if err != nil {
	// 	kcs.logger.WithError(err).Error("Failed to create topic")
	// 	return err
	// }
	// return nil
}

func (kcs *KafkaServerService) DeleteTopic(topic string) error {
	kcs.logger.Info("KafkaServerService DeleteTopic")
	return fmt.Errorf("Not implemented")

	// err := kcs.producer.DeleteTopics([]string{topic}, kafka.SetAdminOperationTimeout(1000))
	// if err != nil {
	// 	kcs.logger.WithError(err).Error("Failed to delete topic")
	// 	return err
	// }
	// return nil
}

func (kcs *KafkaServerService) Produce(message model.MessageDTO) error {
	kcs.logger.Infof("KafkaServerService Produce: %v to Topic %s", message, kcs.topics[0])
	messageBytes, err := json.Marshal(message)
	if err != nil {
		kcs.logger.WithError(err).Error("Failed to marshal message to JSON")
		return err
	}

	err = kcs.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kcs.topics[0], Partition: kafka.PartitionAny},
		Value:          []byte(messageBytes),
	}, nil)

	if err != nil {
		kcs.logger.WithError(err).Error("Failed to produce message")
		return err
	}

	return nil
}

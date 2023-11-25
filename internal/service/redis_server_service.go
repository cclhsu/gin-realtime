package service

// go get -u github.com/go-redis/redis/v8

import (
	"context"
	"encoding/json"
	"fmt"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type RedisServerInterface interface {
	NewRedisServerService(ctx context.Context, logger *logrus.Logger, bootstrapServers string, topics []string) (*RedisServerService, error)
	_interface.ProducerServiceInterface
	Initialize() error
}

type RedisServerService struct {
	ctx      context.Context
	logger   *logrus.Logger
	producer *redis.Client
	topics   []string
}

func NewRedisServerService(ctx context.Context, logger *logrus.Logger, bootstrapServers string, topics []string) (*RedisServerService, error) {
	logger.Info("RedisServerService NewRedisServerService")
	logger.Debugf("RedisServerService NewRedisServerService: %v", bootstrapServers)
	logger.Debugf("RedisServerService NewRedisServerService: %v", topics)

	producer := redis.NewClient(&redis.Options{
		Addr:     bootstrapServers, // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	_, err := producer.Ping().Result()
	if err != nil {
		return nil, err
	}

	logger.Println("Connected to Redis")

	return &RedisServerService{
		ctx:      ctx,
		logger:   logger,
		producer: producer,
		topics:   topics,
	}, nil
}

func (kss *RedisServerService) Initialize() error {
	kss.logger.Info("RedisServerService Initialize")

	// go func() {
	// 	for {
	// 		message, err := kss.producer.BLPop(0, kss.topics...).Result()
	// 		if err != nil {
	// 			kss.logger.WithError(err).Error("Failed to consume message from Redis")
	// 			continue
	// 		}

	// 		kss.logger.Printf("Received task from Redis: %s\n", message[1])

	// 		var messageDTO model.MessageDTO
	// 		err = json.Unmarshal([]byte(message[1]), &messageDTO)
	// 		if err != nil {
	// 			kss.logger.WithError(err).Error("Failed to unmarshal message from JSON")
	// 			continue
	// 		}

	// 		kss.logger.Printf("Received task from Redis: %v\n", messageDTO)
	// 	}
	// }()

	return nil
}

func (kss *RedisServerService) Disconnect() {
	if kss.producer != nil {
		_ = kss.producer.Close()
		kss.logger.Println("Redis client closed")
	}
}

func (kcs *RedisServerService) CreateTopic(topic string) error {
	kcs.logger.Info("RedisServerService CreateTopic")
	return fmt.Errorf("Not implemented")
}

func (kcs *RedisServerService) DeleteTopic(topic string) error {
	kcs.logger.Info("RedisServerService DeleteTopic")
	return fmt.Errorf("Not implemented")
}

func (kcs *RedisServerService) Produce(message model.MessageDTO) error {
	kcs.logger.Infof("RedisServerService Produce: %v to Topic %s", message, kcs.topics[0])
	messageBytes, err := json.Marshal(message)
	if err != nil {
		kcs.logger.WithError(err).Error("Failed to marshal message to JSON")
		return err
	}

	err = kcs.producer.RPush("tasks", string(messageBytes)).Err()
	if err != nil {
		return err
	}

	kcs.logger.Printf("Sent task to Redis: %s\n", string(messageBytes))

	return nil
}

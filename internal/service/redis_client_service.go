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

type RedisClientInterface interface {
	NewRedisClientService(ctx context.Context, logger *logrus.Logger, bootstrapServers, groupID string, topics []string) (*RedisClientService, error)
	_interface.ConsumerServiceInterface
	Initialize() error
	Shutdown() error
}

type RedisClientService struct {
	ctx    context.Context
	logger *logrus.Logger
	// redisServerServiceURL string
	server   string
	GroupID  string
	consumer *redis.Client
	topics   []string
}

func NewRedisClientService(ctx context.Context, logger *logrus.Logger, bootstrapServers, groupID string, topics []string) (*RedisClientService, error) {
	logger.Info("RedisClientService NewRedisClientService")
	logger.Debugf("RedisClientService NewRedisClientService: %v", bootstrapServers)
	logger.Debugf("RedisClientService NewRedisClientService: %v", topics)

	consumer := redis.NewClient(&redis.Options{
		Addr:     bootstrapServers, // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	_, err := consumer.Ping().Result()
	if err != nil {
		return nil, err
	}

	logger.Println("Connected to Redis")

	return &RedisClientService{
		ctx:      ctx,
		logger:   logger,
		server:   bootstrapServers,
		GroupID:  groupID,
		consumer: consumer,
		topics:   topics,
	}, nil
}

func (kcs *RedisClientService) Initialize() error {
	kcs.logger.Info("RedisClientService Initialize")
	go kcs.Consume()
	return nil
}

func (kcs *RedisClientService) Disconnect() {
	if kcs.consumer != nil {
		_ = kcs.consumer.Close()
		kcs.logger.Println("Redis client closed")
	}
}

func (kcs *RedisClientService) Shutdown() error {
	kcs.logger.Info("Shutting down RedisClientService")
	if err := kcs.consumer.Close(); err != nil {
		kcs.logger.WithError(err).Error("Error shutting down Redis consumer")
		return err
	}
	return nil
}

// func (kcs *RedisClientService) initializeRedisServerServiceURL() string {
// 	SERVER_HOST := os.Getenv("SERVER_HOST")
// 	if SERVER_HOST == "" {
// 		SERVER_HOST = "0.0.0.0"
// 	}
// 	SERVER_PORT := os.Getenv("SERVER_PORT")
// 	if SERVER_PORT == "" {
// 		SERVER_PORT = "3001"
// 	}
// 	return fmt.Sprintf("http://%s:%s/redis/handler", SERVER_HOST, SERVER_PORT)
// }

func (kcs *RedisClientService) Consume() error {
	kcs.logger.Info("RedisClientService Consume: %v", kcs.topics)

	for {
		message, err := kcs.consumer.BLPop(0, kcs.topics...).Result()
		if err != nil {
			fmt.Printf("Error consuming messages: %v\n", err)
			return err
		}

		var messageBody model.MessageDTO
		if err := json.Unmarshal([]byte(message[1]), &messageBody); err != nil {
			fmt.Printf("Error unmarshalling task: %v\n", err)
			continue
		}

		fmt.Printf("Received message from Redis: %+v\n", messageBody)
	}
}

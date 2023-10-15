package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type RedisClientInterface interface {
	NewRedisClientService(ctx context.Context, logger *logrus.Logger) *RedisClientService
	Initialize()
	Connection() error
	Disconnection() error
	Trigger(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type RedisClientService struct {
	ctx                   context.Context
	logger                *logrus.Logger
	redisServerServiceURL string
}

func NewRedisClientService(ctx context.Context, logger *logrus.Logger) *RedisClientService {
	return &RedisClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (rcs *RedisClientService) Initialize() {
	rcs.logger.Info("RedisClientService Initialize")
	rcs.redisServerServiceURL = rcs.initializeRedisServerServiceURL()
	rcs.logger.Infof("Redis Server URL: %s\n", rcs.redisServerServiceURL)
}

func (rcs *RedisClientService) initializeRedisServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/ws", SERVER_HOST, SERVER_PORT)
}

func (rcs *RedisClientService) Connection() error {
	rcs.logger.Info("RedisClientService Connection")
	return nil
}

func (rcs *RedisClientService) Disconnection() error {
	rcs.logger.Info("RedisClientService Disconnection")
	return nil
}

func (rcs *RedisClientService) Trigger(message string) (string, error) {
	rcs.logger.Info("RedisClientService Trigger")
	return "", nil
}

func (rcs *RedisClientService) Echo() (string, error) {
	rcs.logger.Info("RedisClientService Echo")
	return "", nil
}

func (rcs *RedisClientService) Broadcast() (string, error) {
	rcs.logger.Info("RedisClientService Broadcast")
	return "", nil
}

func (rcs *RedisClientService) Health() string {
	rcs.logger.Info("RedisClientService Health")
	return "RedisClientService Health"
}

package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cclhsu/gin-realtime/internal/types"
	"github.com/cclhsu/gin-realtime/internal/utils"
)

type ConnectionResponseDTO struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ConnectionDTO struct {
	UUID        string                  `json:"UUID"`
	Type        types.MESSAGE_TYPES     `json:"type"`
	Stage       types.STAGE_TYPES       `json:"stage"`
	Environment types.ENVIRONMENT_TYPES `json:"environment"`
	Sender      string                  `json:"sender"`
	// Timestamp   time.Time               `json:"timestamp"`
}

func ConvertToConnectionResponseDTO(item interface{}) (ConnectionResponseDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return ConnectionResponseDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct ConnectionResponseDTO
	connection := ConnectionResponseDTO{
		Status:  utils.ExtractString(v, "status"),
		Code:    int(utils.ExtractFloat64(v, "code")),
		Message: utils.ExtractString(v, "message"),
		Data:    utils.ExtractMap(v, "data"),
	}

	return connection, nil
}

func ConvertToConnectionDTO(item interface{}) (ConnectionDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return ConnectionDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct ConnectionDTO
	connection := ConnectionDTO{
		UUID:        utils.ExtractString(v, "UUID"),
		Type:        types.MESSAGE_TYPES(utils.ExtractFloat64(v, "type")),
		Stage:       types.STAGE_TYPES(utils.ExtractFloat64(v, "stage")),
		Environment: types.ENVIRONMENT_TYPES(utils.ExtractFloat64(v, "environment")),
		Sender:      utils.ExtractString(v, "sender"),
		// Timestamp:   utils.ExtractTimestamp(v, "timestamp"),
	}

	return connection, nil
}

type HealthResponseDTO struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type HealthDTO struct{}

func ConvertToHealthResponseDTO(item interface{}) (HealthResponseDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return HealthResponseDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct HealthResponseDTO
	health := HealthResponseDTO{
		Status:  utils.ExtractString(v, "status"),
		Code:    int(utils.ExtractFloat64(v, "code")),
		Message: utils.ExtractString(v, "message"),
		Data:    utils.ExtractMap(v, "data"),
	}

	return health, nil
}

func ConvertToHealthDTO(item interface{}) (HealthDTO, error) {
	// Check the type of the item
	_, ok := item.(map[string]interface{})
	if !ok {
		return HealthDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct HealthDTO
	health := HealthDTO{}

	return health, nil
}

type MessageResponseDTO struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type MessageDTO struct {
	UUID          string                  `json:"UUID"`
	Type          types.MESSAGE_TYPES     `json:"type"`
	Action        types.ACTION_TYPES      `json:"action"`
	Stage         types.STAGE_TYPES       `json:"stage"`
	Environment   types.ENVIRONMENT_TYPES `json:"environment"`
	Sender        string                  `json:"sender"`
	Recipient     string                  `json:"recipient"`
	Recipients    []string                `json:"recipients"`
	RecipientType types.RECIPIENT_TYPES   `json:"recipientType"`
	// Timestamp     time.Time               `json:"timestamp"`
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}

func ConvertToMessageResponseDTO(item interface{}) (MessageResponseDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return MessageResponseDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct MessageResponseDTO
	message := MessageResponseDTO{
		Status:  utils.ExtractString(v, "status"),
		Code:    int(utils.ExtractFloat64(v, "code")),
		Message: utils.ExtractString(v, "message"),
		Data:    utils.ExtractMap(v, "data"),
	}

	return message, nil
}

func ConvertToMessageDTO(item interface{}) (MessageDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return MessageDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	recipients, err := utils.ExtractStringArray(v, "recipients")
	if err != nil {
		return MessageDTO{}, err
	}

	// Extract data
	message := MessageDTO{
		UUID:          utils.ExtractString(v, "UUID"),
		Type:          types.MESSAGE_TYPES(utils.ExtractFloat64(v, "type")),
		Action:        types.ACTION_TYPES(utils.ExtractFloat64(v, "action")),
		Stage:         types.STAGE_TYPES(utils.ExtractFloat64(v, "stage")),
		Environment:   types.ENVIRONMENT_TYPES(utils.ExtractFloat64(v, "environment")),
		Sender:        utils.ExtractString(v, "sender"),
		Recipient:     utils.ExtractString(v, "recipient"),
		Recipients:    recipients,
		RecipientType: types.RECIPIENT_TYPES(utils.ExtractFloat64(v, "recipientType")),
		// Timestamp:     utils.ExtractTimestamp(v, "timestamp"),
		Data:     utils.ExtractMap(v, "data"),
		Metadata: utils.ExtractMap(v, "metadata"),
	}

	return message, nil
}

type RegistrationResponseDTO struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type RegistrationDTO struct {
	UUID        string                  `json:"UUID"`
	Type        types.MESSAGE_TYPES     `json:"type"`
	Stage       types.STAGE_TYPES       `json:"stage"`
	Environment types.ENVIRONMENT_TYPES `json:"environment"`
	Sender      string                  `json:"sender"`
	// Timestamp     time.Time                 `json:"timestamp"`
	CallbackURL   string   `json:"callbackURL"`
	Subscriptions []string `json:"subscriptions"`
	// Expires       time.Time                 `json:"expires"`
	Secret string                    `json:"secret"`
	State  types.GENERAL_STATE_TYPES `json:"state"`
}

func ConvertToRegistrationResponseDTO(item interface{}) (RegistrationResponseDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return RegistrationResponseDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct RegistrationResponseDTO
	registration := RegistrationResponseDTO{
		Status:  utils.ExtractString(v, "status"),
		Code:    int(utils.ExtractFloat64(v, "code")),
		Message: utils.ExtractString(v, "message"),
		Data:    utils.ExtractMap(v, "data"),
	}

	return registration, nil
}

func ConvertToRegistrationDTO(item interface{}) (RegistrationDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return RegistrationDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Extract subscriptions
	subscriptions, err := utils.ExtractStringArray(v, "subscriptions")
	if err != nil {
		return RegistrationDTO{}, err
	}

	// Construct RegistrationDTO
	registration := RegistrationDTO{
		UUID:        utils.ExtractString(v, "UUID"),
		Type:        types.MESSAGE_TYPES(utils.ExtractFloat64(v, "type")),
		Stage:       types.STAGE_TYPES(utils.ExtractFloat64(v, "stage")),
		Environment: types.ENVIRONMENT_TYPES(utils.ExtractFloat64(v, "environment")),
		Sender:      utils.ExtractString(v, "sender"),
		// Timestamp:     utils.ExtractTimestamp(v, "timestamp"),
		CallbackURL:   utils.ExtractString(v, "callbackURL"),
		Subscriptions: subscriptions,
		// Expires:       utils.ExtractTimestamp(v, "expires"),
		Secret: utils.ExtractString(v, "secret"),
		State:  types.GENERAL_STATE_TYPES(utils.ExtractFloat64(v, "state")),
	}

	return registration, nil
}

type SubscriptionResponseDTO struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type SubscriptionDTO struct {
	UUID       string                            `json:"UUID"`
	ClientID   string                            `json:"clientId"`
	Topic      string                            `json:"topic"`
	Callback   func(from string, message string) `json:"-"`
	Endpoint   string                            `json:"endpoint"`
	Secret     string                            `json:"secret"`
	Active     bool                              `json:"active"`
	Expiration time.Time                         `json:"expiration"`
}

func ConvertToSubscriptionResponseDTO(item interface{}) (SubscriptionResponseDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return SubscriptionResponseDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct SubscriptionResponseDTO
	subscription := SubscriptionResponseDTO{
		Status:  utils.ExtractString(v, "status"),
		Code:    int(utils.ExtractFloat64(v, "code")),
		Message: utils.ExtractString(v, "message"),
		Data:    utils.ExtractMap(v, "data"),
	}

	return subscription, nil
}

func ConvertToSubscriptionDTO(item interface{}) (SubscriptionDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return SubscriptionDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct SubscriptionDTO
	subscription := SubscriptionDTO{
		UUID:       utils.ExtractString(v, "UUID"),
		ClientID:   utils.ExtractString(v, "clientId"),
		Topic:      utils.ExtractString(v, "topic"),
		Endpoint:   utils.ExtractString(v, "endpoint"),
		Secret:     utils.ExtractString(v, "secret"),
		Active:     utils.ExtractBool(v, "active"),
		Expiration: utils.ExtractTimestamp(v, "expiration"),
	}

	return subscription, nil
}

type TaskResponseDTO struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type TaskDTO struct {
	UUID          string                  `json:"UUID"`
	Type          types.MESSAGE_TYPES     `json:"type"`
	Action        types.ACTION_TYPES      `json:"action"`
	Stage         types.STAGE_TYPES       `json:"stage"`
	Environment   types.ENVIRONMENT_TYPES `json:"environment"`
	Sender        string                  `json:"sender"`
	Recipient     string                  `json:"recipient"`
	Recipients    string                  `json:"recipients"`
	RecipientType types.RECIPIENT_TYPES   `json:"recipientType"`
	// Timestamp     time.Time               `json:"timestamp"`
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}

func ConvertToTaskResponseDTO(item interface{}) (TaskResponseDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return TaskResponseDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Construct TaskResponseDTO
	task := TaskResponseDTO{
		Status:  utils.ExtractString(v, "status"),
		Code:    int(utils.ExtractFloat64(v, "code")),
		Message: utils.ExtractString(v, "message"),
		Data:    utils.ExtractMap(v, "data"),
	}

	return task, nil
}

func ConvertToTaskDTO(item interface{}) (TaskDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return TaskDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Extract data
	task := TaskDTO{
		UUID:          utils.ExtractString(v, "UUID"),
		Type:          types.MESSAGE_TYPES(utils.ExtractFloat64(v, "type")),
		Action:        types.ACTION_TYPES(utils.ExtractFloat64(v, "action")),
		Stage:         types.STAGE_TYPES(utils.ExtractFloat64(v, "stage")),
		Environment:   types.ENVIRONMENT_TYPES(utils.ExtractFloat64(v, "environment")),
		Sender:        utils.ExtractString(v, "sender"),
		Recipient:     utils.ExtractString(v, "recipient"),
		Recipients:    utils.ExtractString(v, "recipients"),
		RecipientType: types.RECIPIENT_TYPES(utils.ExtractFloat64(v, "recipientType")),
		// Timestamp:     utils.ExtractTimestamp(v, "timestamp"),
		Data:     utils.ExtractMap(v, "data"),
		Metadata: utils.ExtractMap(v, "metadata"),
	}
	return task, nil
}

type TopicDTO struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func ConvertToTopicDTO(item interface{}) (TopicDTO, error) {
	// Check the type of the item
	v, ok := item.(map[string]interface{})
	if !ok {
		return TopicDTO{}, fmt.Errorf("Unexpected type in item: %v", reflect.TypeOf(item))
	}

	// Extract data
	topic := TopicDTO{
		Topic:   utils.ExtractString(v, "topic"),
		Message: utils.ExtractString(v, "message"),
	}
	return topic, nil
}

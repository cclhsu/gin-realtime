package model

import "time"

type WebhookType string

const (
	Test	  WebhookType = "test"
	Dev		  WebhookType = "dev"
	Prod	  WebhookType = "prod"
	Broadcast WebhookType = "broadcast"
	System	  WebhookType = "system"
	Exclusive WebhookType = "exclusive"
)

type EventDataResponseDTO struct {
	Success bool		`json:"success"`
	Message string		`json:"message"`
	Data	interface{} `json:"data"`
}

type EventDataDTO struct {
	ID	 string		 `json:"id"`
	Type string		 `json:"type"`
	Data interface{} `json:"data"`
}

type WebhookConfigDTO struct {
	Secret	 string		 `json:"secret"`
	IsActive bool		 `json:"isActive"`
	Type	 WebhookType `json:"type"`
}

type WebhookInfoDTO struct {
	ID		   string			`json:"id"`
	URL		   string			`json:"url"`
	ExpiryDate time.Time		`json:"expiryDate"`
	Config	   WebhookConfigDTO `json:"config"`
}

type WebhookRegistrationResponseDTO struct {
	Success bool		`json:"success"`
	Message string		`json:"message"`
	Data	interface{} `json:"data"`
}

package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type SubscriptionControllerInterface interface {
	Subscribe(ginContext *gin.Context)          // Subscribe to receive messages
	Unsubscribe(ginContext *gin.Context)        // Unsubscribe from receiving messages
	ListSubscriptions(ginContext *gin.Context)  // List subscriptions to receive messages
	UpdateSubscription(ginContext *gin.Context) // Update subscription to receive messages
}

type SubscriptionServiceInterface interface {
	Subscribe(subscription model.SubscriptionDTO) (model.SubscriptionResponseDTO, error)                                 // Subscribe to receive messages
	Unsubscribe(subscriptionID string) error                                                                             // Unsubscribe from receiving messages
	ListSubscriptions() ([]model.SubscriptionDTO, error)                                                                 // List subscriptions to receive messages
	UpdateSubscription(subscriptionID string, subscription model.SubscriptionDTO) (model.SubscriptionResponseDTO, error) // Update subscription to receive messages
}

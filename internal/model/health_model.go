package model

// // HealthService represents the HealthService service.
// type HealthService interface {
//	IsHealthy(*HealthRequest) (*HealthResponse, error)
//	IsALive(*HealthRequest) (*HealthResponse, error)
//	IsReady(*HealthRequest) (*HealthResponse, error)
// }

// ServingStatus represents the enumeration for serving status.
type ServingStatus int32

const (
	ServingStatus_UNKNOWN	  ServingStatus = 0
	ServingStatus_SERVING	  ServingStatus = 1
	ServingStatus_NOT_SERVING ServingStatus = 2
)

// HealthRequest represents the HealthRequest message.
type HealthRequest struct {
	Service string `json:"service"`
}

// HealthResponse represents the HealthResponse message.
type HealthResponse struct {
	Status ServingStatus `json:"status"`
}

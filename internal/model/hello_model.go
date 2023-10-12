package model

// // HelloService represents the HelloService service.
// type HelloService interface {
//	GetHelloString(*wrapperspb.Empty) (*HelloStringResponse, error)
//	GetHelloJson(*wrapperspb.Empty) (*HelloJsonResponse, error)
// }

// HelloStringResponse represents the HelloStringResponse message.
// type HelloStringResponse struct{}
type HelloStringResponse string

// Data represents the Data message.
type Data struct {
	Message string `json:"message"`
}

// HelloJsonResponse represents the HelloJsonResponse message.
type HelloJsonResponse struct {
	Data Data `json:"data"`
}

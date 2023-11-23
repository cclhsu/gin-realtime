package _interface

type GenericClientControllerInterface interface {
	Connect(endpoint string) error								// Connect to a server at the specified endpoint
	Disconnect() error											// Disconnect from the server and stop the client
	CheckHealth() (interface{}, error)							// Check the health of the connection to the server
	SendMessage(to string, message string) error				// Send a directed message to a specific recipient
	Send(data interface{}) error								// Send a message to the server (broadcast)
	Echo(data interface{}) (interface{}, error)					// Send a message and expect to receive the same message
	Receive(callback func(data interface{})) error				// Handle incoming messages from the server
	Subscribe(callback func(from string, message string)) error // Handle incoming messages from the server
	Unsubscribe() error											// Stop handling incoming messages from the server
	Register() error											// Register the client with the server
	Unregister() error											// Unregister the client from the server
	ListConnectedClients() ([]string, error)					// List all connected clients
	ListDisconnectedClients() ([]string, error)					// List all disconnected clients
	ListRegisteredClients() ([]string, error)					// List all registered clients
	ListSubscribedClients() ([]string, error)					// List all subscribed clients
}

type GenericClientServiceInterface interface {
	Initialize() error																	   // Initialize the client
	OnConnected(callback func()) error													   // Handle connection to the server
	OnDisconnected(callback func()) error												   // Handle disconnection from the server
	OnHealthStatus(callback func(status string) (interface{}, error)) (interface{}, error) // Handle health status changes of the connection to the server
	OnMessageReceived(callback func(from string, message string) (interface{}, error))	   // Handle incoming messages from the server
}

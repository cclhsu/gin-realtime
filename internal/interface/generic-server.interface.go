package _interface

type GenericServerControllerInterface interface {
	StartServer(port int) error																	// Start the server on the specified port
	StopServer() error																			// Stop the server and disconnect all clients
	CheckHealth() (interface{}, error)															// Check the health of the server
	HandleClientConnection(client interface{}) (interface{}, error)								// Handle a new client connection
	Broadcast(data interface{}) error															// Send a message to all connected clients (broadcast)
	SendMessage(clientId string, message string) error											// Send a directed message to a specific client (unicast)
	Echo(clientId string, message string) (interface{}, error)									// Send a message and expect to receive the same message
	Subscribe(clientId string, callback func(from string, message string)) (interface{}, error) // Handle incoming messages from a specific client
	Unsubscribe(clientId string) error															// Stop handling incoming messages from a specific client
	Register(clientId string) (interface{}, error)												// Register a client with the server
	Unregister(clientId string) error															// Unregister a client from the server
	ListConnectedClients() ([]string, error)													// List all connected clients
	ListDisconnectedClients() ([]string, error)													// List all disconnected clients
	ListRegisteredClients() ([]string, error)													// List all registered clients
	ListSubscribedClients() ([]string, error)													// List all subscribed clients
}

type GenericServerServiceInterface interface {
	Initialize() error																			  // Initialize the server (start listening for connections)
	OnClientConnected(callback func(client interface{})) (interface{}, error)					  // Handle new client connections
	OnClientDisconnected(callback func(client interface{})) error								  // Handle client disconnections
	OnHealthStatus(callback func(status string)) (interface{}, error)							  // Handle health status changes of the server
	OnMessageReceived(callback func(from string, message string)) (interface{}, error)			  // Handle incoming messages from clients
	OnSubscribe(clientId string, callback func(from string, message string)) (interface{}, error) // Handle incoming messages from a specific client
	OnUnsubscribe(clientId string) error														  // Stop handling incoming messages from a specific client
	OnRegister(clientId string) (interface{}, error)											  // Register a client with the server
	OnUnregister(clientId string) error															  // Unregister a client from the server
}

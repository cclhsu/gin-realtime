package model

// // Client management
// type ClientManager struct {
//	//The client map stores and manages all long connection clients, online is TRUE, and those who are not there are FALSE
//	clients map[*Client]bool
//	//Web side MESSAGE we use Broadcast to receive, and finally distribute it to all clients
//	broadcast chan []byte
//	//Newly created long connection client
//	register chan *Client
//	//Newly canceled long connection client
//	unregister chan *Client
// }

// // Client
// type Client struct {
//	//User ID
//	id string
//	//Connected socket
//	socket *websocket.Conn
//	//Message
//	send chan []byte
// }

// // Will formatting Message into JSON
// type Message struct {
//	//Message Struct
//	Sender	  string `json:"sender,omitempty"`	  //发送者
//	Recipient string `json:"recipient,omitempty"` //接收者
//	Content	  string `json:"content,omitempty"`	  //内容
//	ServerIP  string `json:"serverIp,omitempty"`  //实际不需要 验证k8s
//	SenderIP  string `json:"senderIp,omitempty"`  //实际不需要 验证k8s
// }

// // Create a client manager
// var manager = ClientManager{
//	broadcast:	make(chan []byte),
//	register:	make(chan *Client),
//	unregister: make(chan *Client),
//	clients:	make(map[*Client]bool),
// }

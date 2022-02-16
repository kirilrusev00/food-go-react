package qrdecoderserver

import (
	"net"
)

/*
	ClientManager manages the client connections. It stores them, registers and unregisters them
	and broadcasts the messages to the clients.
*/
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

/*
	Client stores the connection with the client - the socket and the chan to
*/
type Client struct {
	socket net.Conn
	data   chan []byte
}

/*
	start starts the client manager.
*/
func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
			}
		case message := <-manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

/*
	receive is used to receive a message from a client and pass it to broadcast channel in the manager.
*/
func (manager *ClientManager) receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			manager.broadcast <- message
		}
	}
}

/*
	send is used to process the message from a client (a file path) and return the decoded file.
*/
func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			filePath := string(message[:clen(message)])
			decoded := decodeQrCode(filePath)
			client.socket.Write([]byte(decoded))
			client.socket.Write([]byte("\n"))

			manager.unregister <- client
		}
	}
}

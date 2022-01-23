package decoder

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	socket net.Conn
	data   chan []byte
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("One of the connections has terminated!")
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

func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			filePath := string(message[:clen(message)])
			decoded := convertQRCodeToString(filePath)
			client.socket.Write([]byte(decoded))
			client.socket.Write([]byte("\n"))
		}
	}
}

// Remove the unnecessary from the received bytes from the buffer
func clen(n []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] < 32 {
			return i
		}
	}
	return len(n)
}

func convertQRCodeToString(localQrPath string) string {
	result, error := qrDecoder(localQrPath)

	if result == nil || error != nil {
		return "Could not decode QR code"
	}

	return result.String()
}

func qrDecoder(localQrPath string) (*gozxing.Result, error) {
	file, err := os.Open(localQrPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	src := gozxing.NewLuminanceSourceFromImage(img)

	bmp, err := gozxing.NewBinaryBitmap(gozxing.NewGlobalHistgramBinarizer(src))
	if err != nil {
		return nil, err
	}

	qrReader := qrcode.NewQRCodeReader()
	return qrReader.Decode(bmp, nil)
}

func RunServer() {
	fmt.Println("Starting server...")

	network := "tcp"
	address := "localhost:3200"

	listener, error := net.Listen(network, address)
	if error != nil {
		fmt.Println(error)
	}

	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for {
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}
}

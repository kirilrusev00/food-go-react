/*
	Package qrdecoderserver contains functions for the QR Decoder server.
*/
package qrdecoderserver

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net"

	"github.com/kirilrusev00/food-go-react/pkg/config"
)

/*
	QrDecoderServer contains configurations for the QR Decoder server.
*/
type QrDecoderServer struct {
	config   config.QrDecoder
	listener net.Listener
	manager  ClientManager
}

/*
	NewQrDecoderServer creates a new QR Decoder server with the configuration
	variables needed for that.
*/
func NewQrDecoderServer(config config.QrDecoder) (*QrDecoderServer, error) {
	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		return &QrDecoderServer{}, err
	}

	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	qrDecoderServer := &QrDecoderServer{
		config:   config,
		listener: listener,
		manager:  manager,
	}

	return qrDecoderServer, nil
}

/*
	Run starts and runs the QR Decoder server.
*/
func (server *QrDecoderServer) Run() {
	go server.manager.start()
	for {
		connection, err := server.listener.Accept()
		if err != nil {
			log.Println(err)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		server.manager.register <- client
		go server.manager.receive(client)
		go server.manager.send(client)
	}
}

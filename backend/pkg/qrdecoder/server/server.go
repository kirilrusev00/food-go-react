package qrdecoderserver

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net"

	"github.com/kirilrusev00/food-go-react/pkg/config"
)

type QrDecoderServer struct {
	config   config.QrDecoder
	listener net.Listener
	manager  ClientManager
}

func NewQrDecoderServer(config config.QrDecoder) (*QrDecoderServer, error) {
	listener, error := net.Listen("tcp", config.Address)
	if error != nil {
		return &QrDecoderServer{}, error
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

func (server *QrDecoderServer) Run() {
	fmt.Println("Starting server...")

	go server.manager.start()
	for {
		connection, error := server.listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		server.manager.register <- client
		go server.manager.receive(client)
		go server.manager.send(client)
	}
}

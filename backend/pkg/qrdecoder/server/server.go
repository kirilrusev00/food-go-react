package qrdecoderserver

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net"

	"github.com/kirilrusev00/food-go-react/pkg/config"
)

type QrDecoderServer struct {
	config   config.QrDecoder
	listener net.Listener
	manager  ClientManager
}

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

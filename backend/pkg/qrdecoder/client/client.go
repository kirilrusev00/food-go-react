package qrdecoderclient

import (
	"bufio"
	"log"
	"net"

	"github.com/kirilrusev00/food-go-react/pkg/config"
)

type QrDecoderClient struct {
	config      config.QrDecoder
	tmpFilePath string
}

func NewQrDecoderClient(config config.QrDecoder, tmpFilePath string) *QrDecoderClient {
	return &QrDecoderClient{
		config:      config,
		tmpFilePath: tmpFilePath,
	}
}

func (client *QrDecoderClient) ConnectToDecoder() (message string, err error) {
	connType := "tcp"

	log.Println("Connecting to QR Decoder server")

	conn, err := net.Dial(connType, client.config.Address)
	if err != nil {
		log.Println("Error connecting:", err.Error())
		return
	}

	log.Println("Connected to QR Decoder server")

	conn.Write([]byte(client.tmpFilePath))

	message, err = bufio.NewReader(conn).ReadString('\n')

	return
}

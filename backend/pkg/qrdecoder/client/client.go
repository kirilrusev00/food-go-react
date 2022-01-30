package qrdecoderclient

import (
	"bufio"
	"fmt"
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

	fmt.Println("Connecting to", connType, "server", client.config.Address)

	conn, err := net.Dial(connType, client.config.Address)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}

	conn.Write([]byte(client.tmpFilePath))

	message, err = bufio.NewReader(conn).ReadString('\n')

	return
}

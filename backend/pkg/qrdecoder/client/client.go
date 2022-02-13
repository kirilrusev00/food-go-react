/*
	Package qrdecoderclient contains functions for a client to the QR Decoder server.
*/
package qrdecoderclient

import (
	"bufio"
	"log"
	"net"

	"github.com/kirilrusev00/food-go-react/pkg/config"
)

/*
	QrDecoderClient contains configurations for the client to the QR Decoder server.
*/
type QrDecoderClient struct {
	config      config.QrDecoder
	tmpFilePath string
}

/*
	NewQrDecoderClient creates a new client to the QR Decoder server with the configuration
	variables and the path pf a temporary file to be send to the server.
*/
func NewQrDecoderClient(config config.QrDecoder, tmpFilePath string) *QrDecoderClient {
	return &QrDecoderClient{
		config:      config,
		tmpFilePath: tmpFilePath,
	}
}

/*
	ConnectToDecoder is used to connect to the QR Decoder server.
*/
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

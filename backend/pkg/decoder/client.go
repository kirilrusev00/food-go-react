package decoder

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "3200"
	connType = "tcp"
)

func ConnectToDecoder(tmpFilePath string) string {
	fmt.Println("Connecting to", connType, "server", connHost+":"+connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	conn.Write([]byte(tmpFilePath))

	message, _ := bufio.NewReader(conn).ReadString('\n')

	return message
}

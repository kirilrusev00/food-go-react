// Starting point of the Food Analyzer project
package main

import (
	"log"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/database"
	qrdecoderserver "github.com/kirilrusev00/food-go-react/pkg/qrdecoder/server"
	"github.com/kirilrusev00/food-go-react/pkg/server"
)

func main() {
	// Load the environment variables
	config, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("Cannot read env variables:", err)
	}

	// Create and start the QR Decoder server
	qrDecoderServer, err := qrdecoderserver.NewQrDecoderServer(config.QrDecoder)
	if err != nil {
		log.Fatal("Cannot start QR Decoder server:", err)
	}
	go qrDecoderServer.Run()

	log.Println("Started QR Decoder server")

	// Connect to the database
	db, err := database.NewDBConn(config.Database)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	log.Println("Connected to the database")

	// Create and start the server
	server, err := server.NewServer(config, db)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	log.Println("Starting the server...")

	server.Start()
}

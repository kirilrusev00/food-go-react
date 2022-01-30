package main

import (
	"log"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/database"
	qrdecoderserver "github.com/kirilrusev00/food-go-react/pkg/qrdecoder/server"
	"github.com/kirilrusev00/food-go-react/pkg/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot read env variables:", err)
	}

	qrDecoderServer, err := qrdecoderserver.NewQrDecoderServer(config.QrDecoder)
	if err != nil {
		log.Fatal("Cannot start qr decoder server:", err)
	}
	go qrDecoderServer.Run()

	db, err := database.NewDBConn(config.Database)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	server, err := server.NewServer(config, db)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	server.Start()
	// if err != nil {
	// 	log.Fatal("cannot start server:", err)
	// }
}

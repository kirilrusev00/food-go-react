// Package config provides functions for loading environment variables from .env file.
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config contains all of the environment variables.
type Config struct {
	Server    Server
	FoodData  FoodData
	Database  Database
	QrDecoder QrDecoder
}

/*
	Server contains the variables for the server including
	the server address, the address of the client (used for disabling some of the CORS policies)
	and the max file size that can be sent to the server, in MB.
*/
type Server struct {
	Address         string
	ClientAddress   string
	MaxFileSizeInMb int64
}

/*
	FoodData contains the variables for communicating with FoodData Central API
	including the url address of the API and the api key.
*/
type FoodData struct {
	Address string
	ApiKey  string
}

/*
	Database contains the variables for the database including username, password and the address.
*/
type Database struct {
	Username string
	Password string
	Address  string
}

// QrDecoder contains the address of the QR Decoder service.
type QrDecoder struct {
	Address string
}

// LoadConfig loads the environment variables from env file into Config struct.
func LoadConfig(envFilePath string) (config Config, err error) {
	err = godotenv.Load(envFilePath)
	if err != nil {
		return
	}

	config.Server.Address = os.Getenv("SERVER_ADDRESS")
	config.Server.ClientAddress = os.Getenv("CLIENT_ADDRESS")
	config.Server.MaxFileSizeInMb, err = strconv.ParseInt(os.Getenv("MAX_IMAGE_SIZE_IN_MB"), 10, 64)

	config.FoodData.Address = os.Getenv("FOODDATA_CENTRAL_ADDRESS")
	config.FoodData.ApiKey = os.Getenv("FOODDATA_CENTRAL_API_KEY")

	config.Database.Username = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Address = os.Getenv("DB_ADDRESS")

	config.QrDecoder.Address = os.Getenv("QR_DECODER_ADDRESS")

	return
}

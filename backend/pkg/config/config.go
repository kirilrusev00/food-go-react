package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Client    Client
	Server    Server
	FoodData  FoodData
	Database  Database
	QrDecoder QrDecoder
}

type Client struct {
	Address string
}

type Server struct {
	Address         string
	ClientAddress   string
	MaxFileSizeInMb int64
}

type FoodData struct {
	Address string
	ApiKey  string
}

type Database struct {
	Username string
	Password string
	Address  string
}

type QrDecoder struct {
	Address string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return
	}

	config.Client.Address = os.Getenv("CLIENT_ADDRESS")

	config.Server.Address = os.Getenv("SERVER_ADDRESS")
	config.Server.MaxFileSizeInMb, err = strconv.ParseInt(os.Getenv("MAX_IMAGE_SIZE_IN_MB"), 10, 64)

	config.FoodData.Address = os.Getenv("FOODDATA_CENTRAL_ADDRESS")
	config.FoodData.ApiKey = os.Getenv("FOODDATA_CENTRAL_API_KEY")

	config.Database.Username = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Address = os.Getenv("DB_ADDRESS")

	config.QrDecoder.Address = os.Getenv("QR_DECODER_ADDRESS")

	return
}

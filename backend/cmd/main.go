package main

import (
	"github.com/kirilrusev00/food-go-react/pkg/decoder"
	"github.com/kirilrusev00/food-go-react/pkg/server"
)

func main() {
	go decoder.RunServer()

	server.RunServer()
}

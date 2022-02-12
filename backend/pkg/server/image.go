package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kirilrusev00/food-go-react/pkg/fooddata"
	"github.com/kirilrusev00/food-go-react/pkg/models"
	qrdecoderclient "github.com/kirilrusev00/food-go-react/pkg/qrdecoder/client"
)

// The request need to be of Content-Type multipart/form-data
// and have form-data key "photo" and value the image
func (server *Server) createImage(w http.ResponseWriter, request *http.Request) {
	log.Println("Received new request for decoding a QR code")

	err := request.ParseMultipartForm(server.config.Server.MaxFileSizeInMb << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, h, err := request.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newpath := filepath.Join("..", "tmp")
	newErr := os.MkdirAll(newpath, os.ModePerm)
	if newErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpFilePath := "../tmp/" + h.Filename
	tmpfile, err := os.Create(tmpFilePath)
	defer os.Remove(tmpFilePath)
	defer tmpfile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Created new tmp file <%s>", tmpFilePath)

	qrdecoderClient := qrdecoderclient.NewQrDecoderClient(server.config.QrDecoder, tmpFilePath)
	result, err := qrdecoderClient.ConnectToDecoder()

	if len(result) == 0 || result == "Could not decode QR code\n" || err != nil {
		log.Println("There was a problem with qr code decoding")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// There is a newline character in the end that needs to be removed
	result = result[:len(result)-1]

	log.Printf("The decoded qr code is <%s>", result)

	foodsInDatabase := server.dbConn.GetFoodByGtinUpc(result)

	foods := []models.Food{}
	for _, food := range foodsInDatabase {
		foods = append(foods, models.FromFoodModelToFood(food))
	}

	data := models.FoodsJSON{}
	data.Foods = foods

	if len(foods) == 0 {
		log.Printf("Could not find <%s> in the database. Sending request to FoodData Central", result)

		client := fooddata.NewClient(server.config.FoodData)
		data, err = client.GetData(result)
		if err != nil {
			log.Printf("Could not find <%s> in FoodData Central", result)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("Found <%s> in FoodData Central. Adding it to the database", result)

		for _, food := range data.Foods {
			server.dbConn.InsertFood(food)
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Sending response <%s> to client", jsonBytes)

	w.Write(jsonBytes)
}

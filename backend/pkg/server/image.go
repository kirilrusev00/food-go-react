package server

import (
	"encoding/json"
	"io"
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

	qrdecoderClient := qrdecoderclient.NewQrDecoderClient(server.config.QrDecoder, tmpFilePath)
	result, err := qrdecoderClient.ConnectToDecoder()

	if result == "Could not decode QR code\n" || err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// There is a newline character in the end that needs to be removed
	result = result[:len(result)-1]

	foodsInDatabase := server.dbConn.GetFoodByGtinUpc(result)

	foods := []models.Food{}
	for _, food := range foodsInDatabase {
		foods = append(foods, models.FromFoodModelToFood(food))
	}

	data := models.FoodsJSON{}
	data.Foods = foods

	if len(foods) == 0 {
		client := fooddata.NewClient(server.config.FoodData)
		data, err = client.GetData(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, food := range data.Foods {
			server.dbConn.InsertFood(food)
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

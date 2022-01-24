package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"encoding/json"

	"github.com/kirilrusev00/food-go-react/pkg/database"
	"github.com/kirilrusev00/food-go-react/pkg/decoder"
	"github.com/kirilrusev00/food-go-react/pkg/fooddata"
	"github.com/kirilrusev00/food-go-react/pkg/models"
)

const (
	maxSizeFileInMB = 32
)

func RunServer() {
	createImageHandler := http.HandlerFunc(createImage)
	http.Handle("/image", createImageHandler)
	http.ListenAndServe("localhost:3000", nil)
}

// The request need to be of Content-Type multipart/form-data
// and have form-data key "photo" and value the image
func createImage(w http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(maxSizeFileInMB << 20)
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

	result := decoder.ConnectToDecoder(tmpFilePath)

	if result == "Could not decode QR code\n" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// There is a newline character in the end that needs to be removed
	result = result[:len(result)-1]

	foodsInDatabase := database.GetFoodByGtinUpc(result)

	foods := []models.Food{}
	for _, food := range foodsInDatabase {
		foods = append(foods, models.FromFoodModelToFood(food))
	}

	data := models.FoodsJSON{}
	data.Foods = foods

	if len(foods) == 0 {
		data = fooddata.GetData(result)

		for _, food := range data.Foods {
			database.InsertFood(food)
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

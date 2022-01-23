package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kirilrusev00/food-go-react/pkg/decoder"
	"github.com/kirilrusev00/food-go-react/pkg/fooddata"
)

const (
	maxSizeFileInMB = 32
)

func main() {
	go decoder.RunServer()

	createImageHandler := http.HandlerFunc(createImage)
	http.Handle("/image", createImageHandler)
	http.ListenAndServe("localhost:8080", nil)
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

	food := fooddata.GetData(result)

	w.Write(food)
}

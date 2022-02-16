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

/*
	createImage is the handler function for where a request to "/image" endpoint is made
*/
func (server *Server) createImage(w http.ResponseWriter, request *http.Request) {
	log.Println("Received new request for decoding a QR code")

	/*
		Parse the request. It is expected to be of Content-Type multipart/form-data
		and have form-data key "photo" and value the image. The size of the image
		can be max config.Server.MaxFileSizeInMb MBs.
	*/
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

	// The received image is saved in a temporary directory.
	// It will be deleted after sending the response to the client.
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

	// Send the file path to the image to the QR Decoder server and
	// receive the gtinUpc encoded in the image.
	qrdecoderClient := qrdecoderclient.NewQrDecoderClient(server.config.QrDecoder, tmpFilePath)
	result, err := qrdecoderClient.ConnectToDecoder()

	if len(result) == 0 || result == "Could not decode QR code\n" || err != nil {
		log.Println("There was a problem with qr code decoding")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// There is a newline character in the end that needs to be removed.
	result = result[:len(result)-1]

	log.Printf("The decoded qr code is <%s>", result)

	// Search for the gtinUpc in the database
	foodsInDatabase, err := server.dbConn.GetFoodByGtinUpc(result)
	if err != nil {
		log.Printf("There was an error in searching in the database")
	}

	foods := []models.Food{}
	for _, food := range foodsInDatabase {
		foods = append(foods, models.FromFoodModelToFood(food))
	}

	data := models.FoodsJSON{}
	data.Foods = foods

	// If this gtinUpc is not in the database, send a request to FoodData Central API.
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

		// Insert the results into the database.
		for _, food := range data.Foods {
			err := server.dbConn.InsertFood(food)
			if err != nil {
				log.Printf("There was an error in inserting into the database")
			}
		}
	}

	// Convert the result to FoodsJSON model.
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Sending response <%s> to client", jsonBytes)

	// Send the converted result to the client.
	w.Write(jsonBytes)
}

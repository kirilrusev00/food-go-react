/*
	Package fooddata manages the communication with FoodData Central API
	used for retrieving information about foods and storing it in the local database.
*/
package fooddata

import (
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/models"
)

/*
	HttpClient contains the function from net/http package
	that are used for communicating with FoodData Central API.
*/
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

/*
	FoodDataClient contains the configuration variables	and the client
	needed for the communication with FoodData Central API.
*/
type FoodDataClient struct {
	config config.FoodData
	Client HttpClient
}

/*
	NewClient creates a new FoodDataClient.
*/
func NewClient(config config.FoodData) *FoodDataClient {
	foodDataClient := &FoodDataClient{
		config: config,
		Client: &http.Client{},
	}

	return foodDataClient
}

/*
	GetData is used to send a GET request to FoodData Central API.
	A string to serch for should be provided to the function.
	The response is transformed to models.FoodsJSON.
*/
func (foodDataClient *FoodDataClient) GetData(searchInput string) (data models.FoodsJSON, err error) {
	req, err := http.NewRequest("GET", foodDataClient.config.Address, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("generalSearchInput", searchInput)
	q.Add("requireAllWords", "true")
	q.Add("api_key", foodDataClient.config.ApiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := foodDataClient.Client.Do(req)

	if err != nil {
		log.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(resp_body, &data); err != nil {
		log.Println("failed to unmarshal:", err)
		return
	}

	return
}

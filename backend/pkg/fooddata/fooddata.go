package fooddata

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/models"
)

type FoodDataClient struct {
	config config.FoodData
	client *http.Client
}

func NewClient(config config.FoodData) *FoodDataClient {
	foodDataClient := &FoodDataClient{
		config: config,
		client: &http.Client{},
	}

	return foodDataClient
}

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

	resp, err := foodDataClient.client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal(resp_body, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
		return
	}

	return
}

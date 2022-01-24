package fooddata

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/kirilrusev00/food-go-react/pkg/models"
)

func GetData(searchInput string) models.FoodsJSON {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.nal.usda.gov/fdc/v1/search", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("generalSearchInput", searchInput)
	q.Add("requireAllWords", "true")
	q.Add("api_key", "LMm1mjww0SJZFfTe5ie1Dw48cS9jtdxEuI6HhOmf")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return models.FoodsJSON{}
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	var data models.FoodsJSON
	if err := json.Unmarshal(resp_body, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
		return models.FoodsJSON{}
	}

	return data
}

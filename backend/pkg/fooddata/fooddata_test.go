package fooddata_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/fooddata"
	"github.com/kirilrusev00/food-go-react/utils/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	FoodDataClient fooddata.FoodDataClient
)

func init() {
	envFilePath := "../../cmd/.env"
	config, _ := config.LoadConfig(envFilePath)

	FoodDataClient = *fooddata.NewClient(config.FoodData)
	FoodDataClient.Client = &mocks.MockClient{}
}

func TestGetData(t *testing.T) {
	// build response JSON
	json := `{"foods":[{"fdcId":2041155,"description":"RAFFAELLO, ALMOND COCONUT TREAT","gtinUpc":"009800146130","ingredients":"VEGETABLE OILS (PALM AND SHEANUT). DRY COCONUT, SUGAR, ALMONDS, SKIM MILK POWDER, WHEY POWDER (MILK), WHEAT FLOUR, NATURAL AND ARTIFICIAL FLAVORS, LECITHIN AS EMULSIFIER (SOY), SALT, SODIUM BICARBONATE AS LEAVENING AGENT."}]}`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := FoodDataClient.GetData("raffaello treat")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Foods)
	assert.EqualValues(t, 1, len(response.Foods))
	assert.EqualValues(t, 2041155, response.Foods[0].FdcId)
	assert.EqualValues(t, "RAFFAELLO, ALMOND COCONUT TREAT", response.Foods[0].Description)
	assert.EqualValues(t, "009800146130", response.Foods[0].GtinUpc)
	assert.EqualValues(t, "VEGETABLE OILS (PALM AND SHEANUT). DRY COCONUT, SUGAR, ALMONDS, SKIM MILK POWDER, WHEY POWDER (MILK), WHEAT FLOUR, NATURAL AND ARTIFICIAL FLAVORS, LECITHIN AS EMULSIFIER (SOY), SALT, SODIUM BICARBONATE AS LEAVENING AGENT.", response.Foods[0].Ingredients)
}

package database_test

import (
	"testing"

	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/database"
	"github.com/kirilrusev00/food-go-react/pkg/models"
	"github.com/kirilrusev00/food-go-react/utils/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	DbConn database.DbConn
)

func init() {
	envFilePath := "../../cmd/.env"
	config, _ := config.LoadConfig(envFilePath)

	newDbConn, _ := database.NewDBConn(config.Database)
	DbConn = *newDbConn
	// DbConn.Db = &mocks.MockDB{}
}

func TestInsertFood(t *testing.T) {
	// Create a new instance of *MockDB
	mockDB := &mocks.MockDB{}
	DbConn.Db = mockDB

	sampleFood := models.Food{
		FdcId:       2041155,
		Description: "RAFFAELLO, ALMOND COCONUT TREAT",
		GtinUpc:     "009800146130",
		Ingredients: "VEGETABLE OILS (PALM AND SHEANUT). DRY COCONUT, SUGAR, ALMONDS, SKIM MILK POWDER, WHEY POWDER (MILK), WHEAT FLOUR, NATURAL AND ARTIFICIAL FLAVORS, LECITHIN AS EMULSIFIER (SOY), SALT, SODIUM BICARBONATE AS LEAVENING AGENT.",
	}

	// Call the function with mock DB
	DbConn.InsertFood(sampleFood)

	params := mockDB.CalledWith()

	assert.Equal(t, 5, len(params))

	// Normally we shouldn't test equality of query strings
	// since you might miss some formatting quirks like newline
	// and tab characters.
	expectedQuery := `INSERT INTO foods (fdcId, description, gtinUpc, ingredients) VALUES ( ?, ?, ?, ? )`

	// Assert that the first parameter in the call was SQL query
	assert.Equal(t, expectedQuery, params[0])

	assert.Equal(t, sampleFood.FdcId, params[1])
	assert.Equal(t, sampleFood.Description, params[2])
	assert.Equal(t, sampleFood.GtinUpc, params[3])
	assert.Equal(t, sampleFood.Ingredients, params[4])
}

func TestGetFoodByGtinUpc(t *testing.T) {
	// Create a new instance of *MockDB
	mockDB := &mocks.MockDB{}
	DbConn.Db = mockDB

	sampleGtinUpc := "009800146130"

	// Call the function with mock DB
	DbConn.GetFoodByGtinUpc(sampleGtinUpc)

	params := mockDB.CalledWith()

	assert.Equal(t, 2, len(params))

	// Normally we shouldn't test equality of query strings
	// since you might miss some formatting quirks like newline
	// and tab characters.
	expectedQuery := `SELECT * FROM foods WHERE gtinUpc = ?`

	// Assert that the first parameter in the call was SQL query
	assert.Equal(t, expectedQuery, params[0])

	assert.Equal(t, sampleGtinUpc, params[1])
}

func TestGetAllFoods(t *testing.T) {
	// Create a new instance of *MockDB
	mockDB := &mocks.MockDB{}
	DbConn.Db = mockDB

	// Call the function with mock DB
	DbConn.GetAllFoods()

	params := mockDB.CalledWith()

	assert.Equal(t, 1, len(params))

	// Normally we shouldn't test equality of query strings
	// since you might miss some formatting quirks like newline
	// and tab characters.
	expectedQuery := "SELECT * FROM foods"

	// Assert that the first parameter in the call was SQL query
	assert.Equal(t, expectedQuery, params[0])
}

package models_test

import (
	"testing"

	"github.com/kirilrusev00/food-go-react/pkg/models"
	"github.com/stretchr/testify/assert"
)

var (
	foodModel models.FoodModel
)

func init() {
	foodModel.Id = 1
	foodModel.FdcId = 2041155
	foodModel.Description = "RAFFAELLO, ALMOND COCONUT TREAT"
	foodModel.GtinUpc = "009800146130"
	foodModel.Ingredients = "VEGETABLE OILS (PALM AND SHEANUT). DRY COCONUT, SUGAR, ALMONDS, SKIM MILK POWDER, WHEY POWDER (MILK), WHEAT FLOUR, NATURAL AND ARTIFICIAL FLAVORS, LECITHIN AS EMULSIFIER (SOY), SALT, SODIUM BICARBONATE AS LEAVENING AGENT."
}

func TestFromFoodModelToFood(t *testing.T) {
	food := models.FromFoodModelToFood(foodModel)

	assert.EqualValues(t, foodModel.FdcId, food.FdcId)
	assert.EqualValues(t, foodModel.Description, food.Description)
	assert.EqualValues(t, foodModel.GtinUpc, food.GtinUpc)
	assert.EqualValues(t, foodModel.Ingredients, food.Ingredients)
}

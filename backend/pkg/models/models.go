/*
	Package models contains models used in the application.
*/
package models

/*
	FoodsJSON is a model of the response from FoodData Central API.
*/
type FoodsJSON struct {
	Foods []Food `json:"foods"`
}

/*
	Food is part of the model of the response from FoodData Central API.
*/
type Food struct {
	FdcId       int    `json:"fdcId"`
	Description string `json:"description"`
	GtinUpc     string `json:"gtinUpc"`
	Ingredients string `json:"ingredients"`
}

/*
	FoodModel is the model of food in the local database.
*/
type FoodModel struct {
	Id          int    `json:"id"`
	FdcId       int    `json:"fdcId"`
	Description string `json:"description"`
	GtinUpc     string `json:"gtinUpc"`
	Ingredients string `json:"ingredients"`
}

/*
	FromFoodModelToFood is a transform function from FoodModel to Food
*/
func FromFoodModelToFood(foodModel FoodModel) Food {
	food := Food{}

	food.FdcId = foodModel.FdcId
	food.Description = foodModel.Description
	food.GtinUpc = foodModel.GtinUpc
	food.Ingredients = foodModel.Ingredients

	return food
}

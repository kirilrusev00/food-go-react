package models

type FoodsJSON struct {
	Foods []Food `json:"foods"`
}

type Food struct {
	FdcId       int    `json:"fdcId"`
	Description string `json:"description"`
	GtinUpc     string `json:"gtinUpc"`
	Ingredients string `json:"ingredients"`
}

type FoodModel struct {
	Id          int    `json:"id"`
	FdcId       int    `json:"fdcId"`
	Description string `json:"description"`
	GtinUpc     string `json:"gtinUpc"`
	Ingredients string `json:"ingredients"`
}

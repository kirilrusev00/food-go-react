package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kirilrusev00/food-go-react/pkg/models"
)

func dbConn() (db *sql.DB) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "food",
		AllowNativePasswords: true,
		InterpolateParams:    true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertFood(food models.Food) {
	db := dbConn()

	insert, err := db.Query("INSERT INTO foods (fdcId, description, gtinUpc, ingredients) "+
		"VALUES ( ?, ?, ?, ? )", food.FdcId, food.Description, food.GtinUpc, food.Ingredients)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func GetFoodByGtinUpc(gtinUpc string) []models.FoodModel {
	var foods []models.FoodModel
	db := dbConn()

	rows, err := db.Query(`SELECT * FROM foods WHERE gtinUpc = ?`, gtinUpc)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var food models.FoodModel
		err = rows.Scan(&food.Id, &food.FdcId, &food.Description, &food.GtinUpc, &food.Ingredients)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		foods = append(foods, food)
	}

	return foods
}

func GetAllFoods() []models.FoodModel {
	var foods []models.FoodModel
	db := dbConn()

	rows, err := db.Query("SELECT * FROM foods")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var food models.FoodModel

		err = rows.Scan(&food.Id, &food.FdcId, &food.Description, &food.GtinUpc, &food.Ingredients)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		foods = append(foods, food)
	}

	return foods
}

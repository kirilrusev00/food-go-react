package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/models"
)

type DbConn struct {
	db *sql.DB
}

func NewDBConn(config config.Database) (dbConn *DbConn, err error) {
	cfg := mysql.Config{
		User:                 config.Username,
		Passwd:               config.Password,
		Net:                  "tcp",
		Addr:                 config.Address,
		DBName:               "food",
		AllowNativePasswords: true,
		InterpolateParams:    true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	dbConn = &DbConn{
		db: db,
	}
	return
}

func (dbConn *DbConn) InsertFood(food models.Food) {
	insert, err := dbConn.db.Query("INSERT INTO foods (fdcId, description, gtinUpc, ingredients) "+
		"VALUES ( ?, ?, ?, ? )", food.FdcId, food.Description, food.GtinUpc, food.Ingredients)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func (dbConn *DbConn) GetFoodByGtinUpc(gtinUpc string) []models.FoodModel {
	var foods []models.FoodModel

	rows, err := dbConn.db.Query(`SELECT * FROM foods WHERE gtinUpc = ?`, gtinUpc)
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

func (dbConn *DbConn) GetAllFoods(db *sql.DB) []models.FoodModel {
	var foods []models.FoodModel

	rows, err := dbConn.db.Query("SELECT * FROM foods")
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

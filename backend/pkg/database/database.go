/*
	Package database manages the communication with the local foods database.
*/
package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/kirilrusev00/food-go-react/pkg/config"
	"github.com/kirilrusev00/food-go-react/pkg/models"
)

/*
	SqlDB contains the function from github.com/go-sql-driver/mysql package
	that are used for communicating with the database.
*/
type SqlDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// DbConn stores the connection to the local database.
type DbConn struct {
	Db SqlDB
}

/*
	NewDBConn returns a new db connection to a local food database.
*/
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
		Db: db,
	}
	return
}

/*
	InsertFood inserts a new food into the database.
*/
func (dbConn *DbConn) InsertFood(food models.Food) (err error) {
	insert, err := dbConn.Db.Query("INSERT INTO foods (fdcId, description, gtinUpc, ingredients) "+
		"VALUES ( ?, ?, ?, ? )", food.FdcId, food.Description, food.GtinUpc, food.Ingredients)

	if insert != nil {
		defer insert.Close()
	}

	return
}

/*
	GetFoodByGtinUpc returns foods by gtinUpc code.
*/
func (dbConn *DbConn) GetFoodByGtinUpc(gtinUpc string) (foods []models.FoodModel, err error) {
	rows, err := dbConn.Db.Query(`SELECT * FROM foods WHERE gtinUpc = ?`, gtinUpc)
	if rows == nil {
		return
	}

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var food models.FoodModel
		err = rows.Scan(&food.Id, &food.FdcId, &food.Description, &food.GtinUpc, &food.Ingredients)

		if err != nil {
			return
		}

		foods = append(foods, food)
	}

	return
}

/*
	GetAllFoods returns all food in the databse.
*/
func (dbConn *DbConn) GetAllFoods() (foods []models.FoodModel, err error) {
	rows, err := dbConn.Db.Query("SELECT * FROM foods")
	if rows == nil {
		return
	}

	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var food models.FoodModel

		err = rows.Scan(&food.Id, &food.FdcId, &food.Description, &food.GtinUpc, &food.Ingredients)
		if err != nil {
			return
		}

		foods = append(foods, food)
	}

	return
}

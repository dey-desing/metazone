//establecer coneccion con la base de datos

package db

import (
	"fmt"
	"os"
    "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

)

var DB *sql.DB

func Connect() (*sql.DB, error) {

	godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	DB = database
	return database, nil
}
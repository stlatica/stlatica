package inits

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// NewDB initializes DB connection.
func NewDB() {
	dataSourceName := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true`,
		os.Getenv("STLATICA_DB_USER"),
		os.Getenv("STLATICA_DB_PASSWORD"),
		os.Getenv("STLATICA_DB_HOST"),
		os.Getenv("STLATICA_DB_PORT"),
		os.Getenv("STLATICA_DB_NAME"),
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)
}

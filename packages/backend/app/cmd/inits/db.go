package inits

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// NewDB initializes DB connection.
func NewDB() {
	// TODO: port to be read from environment variables https://github.com/stlatica/stlatica/issues/50
	db, err := sql.Open("mysql", "root:@tcp(localhost:4000)/stlatica?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)
}

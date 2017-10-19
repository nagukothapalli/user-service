package backend

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetNewDBConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:nagu123@/gotest")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db

}

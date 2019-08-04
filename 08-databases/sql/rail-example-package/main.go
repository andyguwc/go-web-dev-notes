package main
import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/andyguwc/go-resources/08-databases/sql/rail-example-package/dbutils.1"
)
func main() {
	// Connect to Database
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// Create tables
	dbutils.Initialize(db)
}


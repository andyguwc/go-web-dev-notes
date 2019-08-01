package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


/*
You need to connect to the database before doing anything else.

First declare a variable Db as a pointer to an sql.DB struct, and then use the init function (which is called automatically for every package) to initialize it.

Here the driver registered itself when we imported the driver. As this "github.com/lib/pq" is imported, the init function kicks off and registers itself.
*/ 

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

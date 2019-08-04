package main

import (
      "log"
      "github.com/andyguwc/go-resources/08-databases/postgres/basic-postgres/models"
)


func main() {
  db, err := models.InitDB()
  if err != nil {
    log.Println(db)
  }
}
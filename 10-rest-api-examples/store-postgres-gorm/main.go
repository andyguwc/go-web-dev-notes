/* Implement Ecommerce Store using GORM and Postgres

Inserting JSON into Postgres and retrieve results in the JSON field 

Endpoint Method Description
/v1/user/id GET Get a user using ID
/v1/user POST Create a new user
/v1/user?first_name=NAME GET Get all users by the given first name
/v1/order/id GET Get an order with the given ID
/v1/order POST Create a new order

GORM 
PostgreSQL introduced a new data type for storing the JSON data. PostgreSQL allows users
to insert a jsonb field type, which holds the JSON string.
We will use the JSON field to store and retrieve items in PostgreSQL. For accessing PostgreSQL's JSON store, the normal pq
library is very tedious. So, in order to handle that better, we can use an Object Relational
Mapper (ORM) called GORM.

We replaced the traditional driver with the GORM driver
Used GORM functions for CRUD operations
We inserted JSON into PostgreSQL and retrieved results in the JSON field


The core motto of this project is to show how JSON can be stored and retrieved out of
PostgreSQL. The special thing here is that we queried on the JSON field instead of the
normal fields in the User table

*/


package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
	"github.com/andyguwc/go-resources/10-rest-api-examples/store-postgres-gorm/models"
)

// DB stores the database session imformation. Needs to be initialized once
type DBClient struct {
	db *gorm.DB
}

// UserResponse is the response to be send back for User
type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

// GetUsersByFirstName fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetUsersByFirstName(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	name := r.FormValue("first_name")
	// Handle response details
	var query = "select * from \"user\" where data->>'first_name'=?"
	driver.db.Raw(query, name).Scan(&users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//responseMap := map[string]interface{}{"url": ""}
	respJSON, _ := json.Marshal(users)
	w.Write(respJSON)
}

// GetUser fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	vars := mux.Vars(r)
	// Handle response details
	driver.db.First(&user, vars["id"])
	var userData interface{}
	// Unmarshal JSON string to interface
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//responseMap := map[string]interface{}{"url": ""}
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

// PostUser adds URL to DB and gives back shortened string
func (driver *DBClient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	postBody, _ := ioutil.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/v1/user/{id:[a-zA-Z0-9]*}", dbclient.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", dbclient.PostUser).Methods("POST")
	r.HandleFunc("/v1/user", dbclient.GetUsersByFirstName).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}


/* CURL Commands

curl -X POST \
http://localhost:8000/v1/user \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
"username": "naren",
"email_address": "narenarya@live.com",
"first_name": "Naren",
"last_name": "Arya"
}'


curl -X GET http://localhost:8000/v1/user/1

curl -X GET 'http://localhost:8000/v1/user?first_name=Naren'


*/


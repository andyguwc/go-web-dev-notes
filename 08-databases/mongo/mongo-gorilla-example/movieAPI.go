/* 
Example Using Gorilla mux and Mongo to create a movie API 
With end to end integration of database and HTTP server 
*/

package main 

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB stores the database session information. Needs to be initialized once

// We need this because multiple handler of Mux need this information. Hence we can attach common data to different functions
// we can also just create DB struct as something only containing session 
type DB struct {
	session *mgo.Session
	collection *mgo.Collection
}

type Movie struct {
	ID   bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	Name string         `json:"name" bson:"name"`
	Year string         `json:"year" bson:"year"`
	Directors []string  `json:"directors" bson:"directors"`
	Writers []string `json:"writers" bson:"writers"`
	BoxOffice BoxOffice `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross uint64 `json:"gross" bson:"gross"`
}

// Fetch movie with given ID
func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var movie Movie
	// cannot pass directly the ID to MongoDB because it expects a BSON object
	// we use the bson.ObjectIdHex function. Once we get the ID, that will be loaded into struct object
	err := db.collection.Find(bson.M{"_id":bson.ObjectIdHex(vars["id"])}).One(&movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}


// Post movie adds a new movie to collection 
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	// ready body 
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	//create Hash ID to insert
	movie.ID = bson.NewObjectId()
	err := db.collection.Insert(movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}

// UpdateMovie modifies the data of given resource
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	putBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(putBody, &movie)
	// Create an Hash ID to insert
	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &movie})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Updated succesfully!"))
	}
}

// DeleteMovie modifies the data of given resource
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Deleted succesfully!"))
	}
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	c := session.DB("appdb").C("movies")
	db := &DB{session: session, collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies", db.PostMovie).Methods("POST")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

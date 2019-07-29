// https://medium.com/@johnteckert/building-a-restful-api-with-go-part-1-9e234774b14d

package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"net/http"

	"github.com/gorilla/mux"
)

// Roll is the model for sushi

type Roll struct {
	ID string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name string `json:"name"`
	Ingredients string `json:"ingredients"`
}

var rolls []Roll

// we use the NewEncoder() and Encode() methods to render our rolls slice as json and send it to the response stream. 

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//NewEncoder returns a new encoder that writes to w.
	json.NewEncoder(w).Encode(rolls)
}

// Iterate over all Rolls to pick the item
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	
}


// Post Request 
// First create an instance of our Roll struct, called it newRoll
// Use .NewDecoder() and .Decode() to read data from our requests
// Add ID and have the newly assembled Roll added to Rolls
// Finally send a response back 
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll 
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls)+1)
	rolls = append(rolls, newRoll)
	if err := json.NewEncoder(w).Encode(newRoll); err !=nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
	}

}


func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll 
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
		}

	}

}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(rolls)
}

func main() {

	Roll1 := Roll{
		ID: "1",
		ImageNumber: "8",
		Name: "Spicy Tuna",
		Ingredients: "Tuna, Chili",
	}


	Roll2 := Roll{
		ID: "2",
		ImageNumber: "5",
		Name: "California",
		Ingredients: "Tuna, Avocado",
	}

	rolls = append(rolls, Roll1, Roll2)

	// initialize router
	router := mux.NewRouter()

	// end points
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}


// curl -i localhost:5000/sushi

// curl -X POST -H "Content-Type: application/json" -d '{"ImageNumber": "8", "Name": "Salmon Roll", "Ingredients": "Salmon, Chili"}' http://localhost:5000/sushi
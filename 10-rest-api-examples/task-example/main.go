// https://thenewstack.io/make-a-restful-json-api-go/

package main 

import (
	"fmt"
    "log"
	"net/http"
	"time"
	
	"github.com/gorilla/mux"
)

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}


func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }
	fmt.Fprintln(w, "Todo Index")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos), err != nil {
		panic(err)
	}

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoId := params["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoId)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))

}
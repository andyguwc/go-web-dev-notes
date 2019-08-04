/* 
Example Using Gorilla Mux Handler


*/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}
func main() {
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}


// Custom Paths

// r := mux.NewRouter()
// s := r.PathPrefix("/articles").Subrouter()
// s.HandleFunc("{id}/settings", settingsHandler)
// s.HandleFunc("{id}/details", detailsHandler)

// Path Prefix is a wildcard for matching after a defined path. Uselful for serving static files
// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
// http.FileServer(http.Dir("/tmp/static"))))


// Query based matching 
// r := mux.NewRouter()
// r.HandleFunc("/articles", QueryHandler)
// r.Queries("id", "category") // this limits the query with the preceding URL

// Use request.URL.Query() to obtain query parameters
func QueryHandler(w http.ResponseWriter, r *http.Request){
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"])
}



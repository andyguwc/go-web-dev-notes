/* Build an URL Shortener service 

*/

package main 

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	// "strconv"
	"net/http"

	"github.com/gorilla/mux"

)



// type MyUrl struct {
// 	ID string `json:"id"`
// 	LongUrl string `json:"longurl"`
// 	ShortUrl string `json:"shorturl"`
// }

type Url struct {
	Url string `json:"url"`
}

var UrlMap map[string]string

func createURL(w http.ResponseWriter, r *http.Request) {
	var longUrl Url 
	json.NewDecoder(r.Body).Decode(&longUrl)
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(longUrl.Url)))[:6]
	if _, ok := UrlMap[hash]; ok {
		w.WriteHeader(200)
	} else {

		w.Header().Set("Content-Type", "application/json")
		UrlMap[hash] = longUrl.Url
		fmt.Println(UrlMap)
	}
}


func redirectURL(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	shortUrl := params["url"]
	fmt.Println(shortUrl)
	// w.Header().Set("Content-Type","application/json")
	// longUrl, _ := UrlMap[shortUrl]
	// fmt.Println(longUrl)

	if longUrl, ok := UrlMap[shortUrl]; ok {
		http.Redirect(w, r, longUrl, 301)
		fmt.Println(longUrl)
	}
}

func main() {
	UrlMap = make(map[string]string)
	router := mux.NewRouter()
	router.HandleFunc("/api/new", createURL).Methods("POST")
	router.HandleFunc("/api/{url}", redirectURL).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// curl -X POST -H "Content-Type: application/json" -d '{"url": "www.google.com"}' http://localhost:8080/api/new
// curl -i http://localhost:8080/api/191347
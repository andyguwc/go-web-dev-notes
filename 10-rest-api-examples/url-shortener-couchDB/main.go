package main 

import (

	"encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/couchbase/gocb"
    "github.com/gorilla/mux"
    "github.com/speps/go-hashids"

)

type MyUrl struct {
	ID string `json:"id"`
	LongUrl string `json:"longurl"`
	ShortUrl string `json:"shorturl"`
}

var bucket *gocb.bucket
var bucketName string 

func main() {
	router := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucketName = "example"
	bucket, _ = cluster.OpenBucket(bucketName, "")
	router.HandleFunc("/{id}", RootHandler).Methods("GET")
	router.HandleFunc("/expand/", ExpandHandler).Methods("GET")
	router.HandleFunc("/create", CreateHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":12345", router))

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var url MyUrl
	_ = json.NewDecoder(r.Body).Decode(&url)
	var n1qlParams []interface{}
        n1qlParams = append(n1qlParams, url.LongUrl)
        query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE longUrl = $1")
		rows, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return 
	}
}

func ExpandHandler(w http.ResponseWriter, r *http.Request) {
	var n1qlParams []interface{}
        query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE shortUrl = $1")
        params := r.URL.Query()
        n1qlParams = append(n1qlParams, params.Get("shortUrl"))
        rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
        var row MyUrl
        rows.One(&row)
        json.NewEncoder(w).Encode(row)

}

func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var url MyUrl
	bucket.Get(params["id"], &url)
	http.Redirect(w, req, url.LongUrl, 301)
}
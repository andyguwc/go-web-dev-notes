// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// https://github.com/TutorialEdge/create-rest-api-in-go-tutorial/blob/master/main.go

package main 


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func homePage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func returnAllArticles (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get All Articles")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Header().Set("Content-Type","application/json")

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
	// fmt.Fprintf(w, "Return single article whose id is: %v\n", id)
}

func createNewArticle (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create article ")
	var newArticle Article 
	json.NewDecoder(r.Body).Decode(&newArticle)
	Articles = append(Articles, newArticle)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(newArticle)
}


func deleteArticle (w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type","application/json")
	fmt.Fprintf(w, "delete article")
	id := mux.Vars(r)["id"]
	for i, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:i],Articles[i+1:]...)
		}
	}
	// json.NewEncoder(w).Encode(Articles)

}

func main() {
	// get some dummy data 
	Articles = []Article{
        Article{Id:"1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id:"2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllArticles).Methods("GET")
	router.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	router.HandleFunc("/article", createNewArticle).Methods("POST")
    router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))

}

// curl -i localhost:8081/all
// curl -X DELETE localhost:8081/article/1
// curl -X POST -H "Content-Type: application/json" -d '{"id":"3", "title":"Hello 3", "desc":"Article Description", "content":"Article Content"}' http://localhost:8081/article


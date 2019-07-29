/*
Serving a subdirectory 

*/

package main 

import (
	"net/http"
)
func main() {
	dir := http.Dir("./files/")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static", handler)

	// http.HandleFunc("/", homePage) // serves a home page 
	http.ListenAndServe(":80080", nil)
}


// using path package path resolution

// func main() {
// 	pr := newPathResolver()
// 	pr.Add("GET /hello", hello)
// 	dir := http.Dir("./files")
// 	handler := http.StripPrefix("/static/", http.FileServer(dir))
// 	pr.Add("GET /static/*", handler.ServeHTTP)
// 	http.ListenAndServe(":8080", pr)
// }


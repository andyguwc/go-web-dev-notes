// Serve file with custom handler
// package main 

// import (
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", readme)
// 	http.ListenAndServe(":8080", nil)
// }

// func readme(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./files/readme.txt") // serve the content of a readme file 
// }



// Use the FileServer handler from http package

package main

import (
	"net/http"
)
func main() {
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", http.FileServer(dir))
}



// # Serving files

// # ServeContent
// [http.ServeContent](https://godoc.org/net/http#ServeContent)
// ``` go
// func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
// ```

// ***

// # ServeFile

// [http.ServeFile](https://godoc.org/net/http#ServeFile)
// ``` go
// func ServeFile(w ResponseWriter, r *Request, name string)
// ```
// ***

// # FileServer & StripPrefix

// [http.FileServer](https://godoc.org/net/http#FileServer)
// ``` go
// func FileServer(root FileSystem) Handler
// ```

// [http.StripPrefix](https://godoc.org/net/http#StripPrefix)
// ``` go
// func StripPrefix(prefix string, h Handler) Handler
// ```
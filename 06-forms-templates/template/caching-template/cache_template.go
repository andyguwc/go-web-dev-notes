/*
instead of parsing template in handler so have to parse each time for a response
parse it outside and reuse the parsed template
*/

package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("./templates/simple.html"))

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

// An input field handling multiple files needs to have only the multiple attribute on it.
// <input type="file" name="files" id="files" multiple>
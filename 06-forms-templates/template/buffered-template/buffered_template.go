/*
Buffer first to make sure no errors, then write the output to the end users

*/

package main 

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./templates/simple.html"))
}

type Page struct {
	Title, Content string 
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p :=&Page{
		Title: "An example",
		Content: "Example content",
	}

	var b bytes.Buffer

	err := t.Execute(&b, p) // buffer to store output of the executed template
	if err != nil {
		fmt.Fprint(w, "Error occured")
		return
	}

	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
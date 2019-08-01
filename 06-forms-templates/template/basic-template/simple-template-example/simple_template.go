/*
Pass simple template

*/

package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}

	// parse a template for later use 
	t := template.Must(template.ParseFiles("simple.html"))

	// write to http output using template and dataset 
	// data will be escaped 
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}


// Parsing and Execute Templates

// func (t *Template) Parse(text string) (*Template, error)
// func (t *Template) ParseFiles(filenames ...string) (*Template, error)

// func (t *Template) Execute(wr io.Writer, data interface{}) error 
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error // execute specific templates 

// *Template is just a container where all the templates are parsed and held 
// excute implements the writer interface 



/*
Render an object to HTML and pass the rendered object to a higher level template

For example render a user object instance on its own and pass the rendered content to a higher level template
*/

package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML // variable to hold data shared between requests

func init() {
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Page struct {
	Title   string
	Content template.HTML
}

type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote: `You keep using that word. I do not think
				it means what you think it means.`,
		Person: "Inigo Montoya",
	} // populates a dataset to supply to template
	var b bytes.Buffer // writes template and adata
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String()) // store quotes as html in global variable

	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "A Quote",
		Content: qc,
	} // create page dataset with quote HTML
	t.ExecuteTemplate(w, "index.html", p) // write quote and page to web server output
}



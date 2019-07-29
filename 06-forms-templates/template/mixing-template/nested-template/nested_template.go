/*
Nesting templates

Avoid common sections of HTML markup in each template 

When the parent is executed to render the output, the subtemplates are included as well 

On go side, just need to make sure to load in two templates

https://godoc.org/text/template#hdr-Nested_template_definitions
*/
package main 
import (
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.html", "head.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An example",
		Content: "Content example",
	}

	t.ExecuteTemplate(w, "index.html", p)

}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

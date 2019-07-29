/* Example of extending template with own functionality 
make your go functions available in templates 
go actions, data, and commands encloded in {{}} can have commands that acts on the data. These commands can be chained into pipelines separated by |


Making custom functions available in templates requires two steps. 
First, a map needs to be created in which names to be used inside the template are mapped to functions in Go. Here dateFormat is mapped to the name dateFormat.
After new template.Template instance is created, the function map (funcMap) needs to be passed into Funcs to make the function mapping available to new templates
Parse the tempalte into template.Template

*/


package main

import (
	"html/template"
	"net/http"
	"time"
)

// html template as string 

// pipe dates through the dateformat command 
var tpl = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Date Example</title>
  </head>
  <body>
  	<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
  </body>
</html>`

// map go functions to template functions 
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

// go function to format time 
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	// create a new template.Template instance
	t := template.New("date")

	// pass additional functions in map into template engine
	t.Funcs(funcMap)

	// parse template string into the template engine
	t.Parse(tpl)
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	// send template with data to output response
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}

// go function to create templates and add template functions each time 
// func parseTemplateString(name, tpl string) *template.Template {
// 	t:= template.New(name)
// 	t.Funcs(funcMap)
// 	t = template.Must(t.Parse(tpl))
// 	return t
// }
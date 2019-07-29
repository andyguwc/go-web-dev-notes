/*
Serving files from alternative locations

Clean separation of production vs. testing environments

In each environment, you’ll want to have a copy or a representative copy of the application’s
assets. When the location of the files is different in each environment, the location needs to
be passed into the application as configuration. This can happen via a shared configuration
service such as etcd, in configuration files, as arguments passed into the application
at startup time, or some other means of passing configuration.

This example pass in location from command line 
*/


package main

import (
	"flag"
	"html/template"
	"net/http"
)

var t *template.Template
var l = flag.String("location", "http://localhost:8080", "A location.") // gets the location of the static files from the application arguments

var tpl = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>A Demo</title>
    <link rel="stylesheet" href="{{.Location}}/styles.css">
  </head>
  <body>
  	<p>A demo.</p>
  </body>
</html>`

func init() {
	t = template.Must(template.New("date").Parse(tpl))
}

func servePage(res http.ResponseWriter, req *http.Request) { // HTTP handler passing the location into the template
	data := struct{ Location *string }{
		Location: l,
	}
	t.Execute(res, data)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}


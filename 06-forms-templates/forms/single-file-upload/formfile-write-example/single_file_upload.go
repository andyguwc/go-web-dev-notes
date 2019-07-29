/*
Upload a file via a form

Being multipart allows the text part of the form to be uploaded and procesed as text, while the file is handles using itw own type

This handler, meant to be used with the web server in the http package, handles both
displaying the form and processing the submitted form.

When a GET request is submitted, it returns the form. When another HTTP method is used, such as a POST or PUT request, the form
submission is processed.
*/


package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {

	// when path is accessed with a GET request, displays the HTML page and form
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
	} else {
		// get the file handler, header, and error for the form field (name)
		f, h, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// create a local location to save the file, production environment would be a file store location
		filename := "/tmp/" + h.Filename
		out, _ := os.Create(filename)
		defer out.Close()

		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}

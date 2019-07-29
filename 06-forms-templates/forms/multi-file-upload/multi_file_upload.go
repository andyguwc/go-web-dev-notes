/*
Multiple value file upload 

After the form has been parsed, the fields are available on MultipartForm. The
uploads to the file-input field with the name files are available on the File property
of MultipartForm as a slice of values. Each value is a *multipart.FileHeader object.

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
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_multiple.html")
		t.Execute(w, nil)
	} else {
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		// get slice containing the files from the MultipartForm
		data := r.MultipartForm
		files := data.File["files"]
		for _, fh := range files {
			f, err := fh.Open()
			defer f.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			// Create files to store the contents of the uploaded file 
			out, err := os.Create("/tmp/" + fh.Filename)
			defer out.Close()
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			// Copies the uploaded file to the location on the filesystem
			_, err = io.Copy(out, f)

			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}


// make sure uploading the right type of files

// file, header, err := r.FormFile("file")
// contentType := header.Header["Content-Type"][0]

// or using the file extension
// file, header, err := r.FormFile("file")
// extension := filepath.Ext(header.Filename)
// type := mime.TypeByExtension(extension)

// or parse the file and detect file type
// file, header, err := r.FormFile("file")
// buffer := make([]byte, 512)
// _, err = file.Read(buffer)
// filetype := http.DetectContentType(buffer)


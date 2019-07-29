/*
Incrementally save uploaded file 

Instead of using API server as a pass through for data 

The way to access the multipart stream directly, which is what ParseMultipartForm
does, is to retrieve the reader from the Request with MultipartReader. After you have
the reader, you can loop over the parts and read each one as it comes in.


*/



package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file_plus.html")
		t.Execute(w, nil)
	} else {

		// Retrieves the multipart reader giving access to the uploaded files and handles any errors
		mr, err := r.MultipartReader()
		values := make(map[string][]string)

		if err != nil {
			panic("Failed to read multipart message")
		}

		maxValueBytes := int64(10 << 20)
		for {
			// Continues looping until all of the multipart message has been read
			// breaking the loop if the end of  request is reached 
			part, err := mr.NextPart()
			if err == io.EOF {

				break
			}

			// retrieves name of the form field
			name := part.FormName()
			if name == "" {
				continue
			}

			// retrieves name of the file if one exists
			filename := part.FileName()

			// buffer to read value of text fields to 
			var b bytes.Buffer
			if filename == "" {
				n, err := io.CopyN(&b, part, maxValueBytes)
				if err != nil && err != io.EOF {
					fmt.Fprint(w, "Error processing form")
					return
				}
				maxValueBytes -= n
				if maxValueBytes == 0 {
					fmt.Fprint(w, "multipart message too large")
					return
				}
				values[name] = append(values[name], b.String())
				continue
			}

			// create location for the filesystem to store the content of a file
			dst, err := os.Create("/tmp/dstfile." + filename)
			defer dst.Close()
			if err != nil {
				return
			}

			// As the file content of a part is uploaded, writes it to the file
			for {
				buffer := make([]byte, 100000)
				cBytes, err := part.Read(buffer)
				if err == io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Println("Upload done")
		fmt.Println(values)

		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}
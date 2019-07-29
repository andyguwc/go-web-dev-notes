# HTML Forms

Within the <form> tag, we place a number of HTML form elements including text input, text area, radio buttons, checkboxes, and file uploads. These elements allow users to enter data to be submitted to the server. Data is submitted to the server when the user clicks a button or somehow triggers the form submission.

HTML Form 
<form action="/process" method="post">
    <input type="text" name="first_name"/>
    <input type="text" name="last_name"/>
    <input type="submit"/>
</form>


The format of the name-value pairs sent through a POST request is specified by the content type of the HTML form. This is defined using the enctype attribute
<form action="/process" method="post" enctype="application/x-www-form-urlencoded">

If you set enctype to multipart/form-data, each name-value pair will be converted
into a MIME message part, each with its own content type and content disposition.

If you’re sending simple text data, the URL encoded form is better—it’s simpler and more efficient and less processing is needed.
If you’re sending large amounts of data, such as uploading files, the multipart-MIME
form is better


# Form Data

## Form
The functions in Request that allow us to extract data from the URL and/or the
body revolve around the Form, PostForm, and MultipartForm fields.

1 Call ParseForm or ParseMultipartForm to parse the request.
2 Access the Form, PostForm, or MultipartForm field accordingly.

package main 
import (
    "fmt"
    "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
    r.ParseForm
    fmt.Fprintln(w, r.Form)
}

func main() {
    server := http.Server(
        Addr: "127.0.0.1:8080",
    )
    http.HandleFunc("/process", process)
    server.ListenAndServe()
}

The Form struct is a map, whose keys are strings and values are a slice
of strings. Notice that the map isn’t sorted, so you might get a different sorting of the
returned values.

## PostForm

Need just the form key-value pairs and want to totally ignore the URL key-value pairs


## MultipartForm
r.ParseMultipartForm(1024)
fmt.Fprintln(w, r.MultipartForm)


## Using FormValue
No need to call ParseForm or ParseMultiForm. It only returns the first value though 
fmt.Fprintln(w, r.FormValue("hello"))

PostFormValue does the same thing for post forms


## Files 


As you can see, you no longer have to call the ParseMultipartForm method, and the
FormFile method returns both the file and the file header at the same time. You simply
need to process the file that’s returned.

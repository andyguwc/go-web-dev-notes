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


# Templates 
Template engines often combine data with templates to produce the final HTML. Handlers usually call template engines to combine data with the templates and return the resultant HTML to the client.

The biggest argument against embedded logic template engines is that presentation and logic are mixed up together and logic is distributed in multiple places, resulting in code that’s hard to maintain. 

The counter-argument against logic-less template engines is that the ideal logic-less template engine would be impractical and that placing more logic into the handlers, especially for presentation, would add unnecessary complexity to the handlers.


## Using Web Template
1 Parse the text-formatted template source, which can be a string or from a template
file, to create a parsed template struct.
2 Execute the parsed template, passing a ResponseWriter and some data to it.
This triggers the template engine to combine the parsed template with the data
to generate the final HTML that’s passed to the ResponseWriter.


ParseFiles vs. ParseGlob
ParseGlob uses pattern matching instead of specific files
t, _ := template.ParseFiles("tmpl.html")
t, _ := template.ParseGlob("*.html")


t := template.Must(template.ParseFiles("tmpl.html"))

executes the first template
t.Execute(w, "Hello World!")

execute specific template
t.ExecuteTemplate(w, "t2.html", "Hello World!")

## Actions

- Conditional actions
- Iterator actions
- Set actions
- Include actions

Conditional action
```
{{ if arg }}
some content
{{ end }}
```

The other variant is
```
{{ if arg }}
some content
{{ else }}
other content
{{ end }}
```

Iterator Action 
```
{{ range array }}
Dot is set to the element {{ . }}
{{ end }}
```

Example range template 
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>Go Web Programming</title>
</head>
<body>
<ul>
{{ range . }}
<li>{{ . }}</li>
{{ end}}
</ul>
</body>
</html>

func process(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl.html")
    daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    t.Execute(w, daysOfWeek)
}

Set action
```
{{ with arg }}
Dot is set to arg
{{ end }}
```

Include action

```
{{ template "t2.html" }}
```

t1 template with an argument passed to t2
```
<html>
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=9">
<title>Go Web Programming</title>
</head>
<body>
<div> This is t1.html before</div>
<div>This is the value of the dot in t1.html - [{{ . }}]</div>
<hr/>
{{ template "t2.html" . }}
<hr/>
<div> This is t1.html after</div>
</body>
</html>
```


Variables 

{{ range $key, $value := . }}
The key is {{ $key }} and the value is {{ $value }}
{{ end }}

here (.) is a map and range initializes the $key and $value variables with the key and value of the successive elements in the map

Pipelines
{{ p1 | p2 | p3 }}
Pipelines are arguments, functions, and methods chained together in a sequence


Functions
Create a FuncMap map (name of function as key and actual function as value) 
Attach the FuncMap to the template


Go also escapes templates based on the context (HTML or Javascript)


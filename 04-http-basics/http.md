# HTTP Request
- request line
- request headers
- optional message body

Example request:

GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
(empty line)

## Request Struct 
Important Parts of Request are: 
- URL
- Header
- Body
- Form, PostForm, and MultipartForm

type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string


    // This field is only available after ParseForm is called.
    Form url.Values

    // This field is only available after ParseForm is called.
    PostForm url.Values

    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. 
    RemoteAddr string
}

## Request URL
URL field is a pointer to the url.URL type
URL *url.URL

type URL struct {
    Scheme string
    Opaque string
    User *Userinfo
    Host string
    Path string
    RawQuery string
    Fragment string
}

scheme://[userinfo@]host/path[?query][#fragment]

## Request Header 

To get one particular header use options below 
h := r.Header["Accept-Encoding"]
h := r.Header.Get("Accept-Encoding")

request headers:
Request headers are colon-separated name-value pairs in plain text, terminated by
a carriage return (CR) and line feed (LF).
- Accept text/html signals to the server that the client wants the response body’s content type to be in HTML
- Accept Charset: Accept-Charset: utf-8 tells the server that the client wants the response body to be in UTF-8.
- Authorization: send Basic Authentication credentials to the server
- Cookie: send back cookies set by the calling server. Cookie: my_first_cookie=hello; my_second_cookie=world
- Content-Length: length of request body
- Content-Type: content type of request body - when uploading file, use multipart/form-data
- Host: name of server and port (default port 80)
- Referrer: address of previous page
- User-Agent: describes the calling agent

request methods: 
- GET tells the server to return the specified resource
- POST data in the message body should be passed to the resource identified by URI
- PUT data in the message body should be the resource at URI. Data is replaced 
- DELETE remove the data identified by URI
- HEAD tells server to not return a message body 
a method is considered safe if it doesn't change the state of the server 

## Request Body

func body(w http.ResponseWriter, r *http.Request) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    fmt.Fprintln(w, string(body))
}


# Request Parsing Form Data 
Retrieve URL & Form data
http.Request is a struct. It has the fields Form & PostForm

## Form
Form contains the parsed form data, including both the URL field's query parameters and the POST or PUT form data.
This field is only available after **ParseForm** is called.
The HTTP client ignores Form and uses Body instead.
Form url.Values

## PostForm
PostForm contains the parsed form data from POST, PATCH, or PUT body parameters.
This field is only available after **ParseForm** is called.
The HTTP client ignores PostForm and uses Body instead. PostForm url.Values

## MultipartForm


## FormValue
FormValue returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. FormValue calls ParseMultipartForm and ParseForm if necessary and ignores any errors returned by these functions. If key is not present, FormValue returns the empty string. To access multiple values of the same key, call ParseForm and then inspect Request.Form directly.



# HTTP Response
status line
headers
optional message body

example response
200 OK
Date: Sat, 22 Nov 2014 12:58:58 GMT
Server: Apache/2
Last-Modified: Thu, 28 Aug 2014 21:01:33 GMT
Content-Length: 33115
Content-Type: text/html; charset=iso-8859-1
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/
TR/xhtml1/DTD/xhtml1-strict.dtd"> <html xmlns='http://www.w3.org/1999/
xhtml'> <head><title>Hypertext Transfer Protocol -- HTTP/1.1</title></
head><body>…</body></html>

Status code:
1xx: informational 
2xx: has processed successfully
3xx: redirection
4xx: client error something wrong with the request
5xx: server error 500 Internal Server Error 


response headers:
- Allow tells client which requests methods are supported by the server
- Content-Length
- Content-Type content type of the response
- Date current time
- Server domain name of server that's returning response
- Set-Cookie set a cookie at the client 


URI
- scheme name
- hierarchical part
- query
- fragment 

Let’s look at an example of an HTTP scheme URI: http://sausheong:password
@www.example.com/docs/file?name=sausheong&location=singapore#summary
The scheme is http, followed by the colon. The segment sausheong:password followed
by the at sign (@) is the user and password information. This is followed by the
rest of the hierarchical part, www.example.com/docs/file.

## ResponseWriter Interface 

The ResponseWriter interface has three methods:
- Write
- WriteHeader
- Header

Write
Example writing HTML string to the HTTP response body using ResponseWriter 
```
func writeExample(w http.ResponseWriter, r *http.Request) {
    str := `<html>
        <head><title>Go Web Programming</title></head>
        <body><h1>Hello World</h1></body>
        </html>`
        w.Write([]byte(str))
}
```

WriteHeader
The WriteHeader method’s name is a bit misleading. It doesn’t allow you to write any headers (you use Header for that), but it takes an integer that represents the status code of the HTTP response and writes it as the return status code for the HTTP response.

```
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(501)
    fmt.Fprintln(w, "No such service, try next door")
}
```

Header 
Finally the Header method returns a map of headers that you can modify. The modified headers will be in the HTTP response that’s sent to the client.

```
func headerExample(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "http://google.com")
    w.WriteHeader(302)
}
```

Example returning JSON to client 
```
func jsonExample(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    post := &Post{
        User: "Sau Sheong",
        Threads: []string{"first", "second", "third"},
    }
    json, _ := json.Marshal(post)
    w.Write(json)
}
```


# Handler 
A handler receives and processes the HTTP request sent from the client. It also calls the
template engine to generate the HTML and finally bundles data into the HTTP
response to be sent back to the client.
https://github.com/GoesToEleven/golang-web-dev/tree/master/017_understanding-net-http-package

Handler 
http.Handler

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

Server
http.ListenAndServe

func ListenAndServe(addr string, handler Handler) error
http.ListenAndServeTLS

func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error


# Web Application
- take input through HTTP from the client 
- process HTTP request message
- generate HTMl and return it in an HTTP response message


MVC model
- The model is a representation of the underlying data, 
- the view is a visualization of the model for the user, and 
- the controller uses input from the user to modify the model. When the model changes, the view updates automatically to show the latest visualization.


Template engine
A template is code that can be converted into HTML that’s sent back to the client in an
HTTP response message


Receiving and processing requests 
1 A client sends a request to a URL at the server.
2 The server has a multiplexer, which redirects the request to the correct handler
to process the request.
3 The handler processes the request and performs the necessary work.
4 The handler calls the template engine to generate the correct HTML to send
back to the client.


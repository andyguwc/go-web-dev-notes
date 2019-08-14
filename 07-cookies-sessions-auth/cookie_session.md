# Cookies

A cookie is a small piece of information that’s stored at the client, originally sent
from the server through an HTTP response message. Every time the client sends an
HTTP request to the server, the cookie is sent along with it. Cookies are designed to
overcome the stateless-ness of HTTP.

Cookie struct

type Cookie struct {
    Name string
    Value string
    Path string
    Domain string
    Expires time.Time
    RawExpires string
    MaxAge int
    Secure bool
    HttpOnly bool
    Raw string
    Unparsed []string
}

Expires was deprecated in favor of MaxAge in HTTP 1.1

Sending cookies to the browser using SetCookie 

func setCookie(w http.ResponseWriter, r *http.Request) {
    c1 := http.Cookie{
        Name: "first_cookie",
        Value: "Go Web Programming",
        HttpOnly: true,
    }
    c2 := http.Cookie{
        Name: "second_cookie",
        Value: "Manning Publications Co",
        HttpOnly: true,
    }
    http.SetCookie(w, &c1)
    http.SetCookie(w, &c2)
}


Get cookie

func getCookie(w http.ResponseWriter, r *http.Request) {
    c1, err := r.Cookie("first_cookie")
    if err != nil {
        fmt.Fprintln(w, "Cannot get the first cookie")
    }
    cs := r.Cookies()
    fmt.Fprintln(w, c1)
    fmt.Fprintln(w, cs)
}


# Session Based Authentication
A client (for example, a browser) sends a request to the Login API of the server. The server tries to check those credentials with the database
and if credentials exist, writes a cookie back onto the response saying this user is authenticated. A cookie is a message to be consumed by the server at the later point of time.

When the client receives the response, it stores that cookie locally. If the web browser is the client, it stores it in the cookie storage. From next time, the client can freely ask for resources from the server by showing the cookie as the key for passage. When a client decides to terminate the session, it calls the Logout API on the server. The server destroys the session in the response. This process continues. The server can also keep an expiration on cookies so that the authentication window is valid for a certain time if there is no activity. This is how all websites work.


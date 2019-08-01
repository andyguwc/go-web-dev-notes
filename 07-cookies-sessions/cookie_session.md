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


/*
Cookies are only sent by the browser to the domain which wrote them.

With every request to a specific domain, the client's web browser looks to see if there is a cookie from that domain on the client's machine. If there is a cookie that has been written by that particular domain, then the browser will send the cookie with every request to that domain.

*/



package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// default route set cookie

func set(w http.ResponseWriter, req *http.Request) {

	// composite literal using name value fields of a cookie 
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path: "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

// read runs the cookie method which takes a name and give the value of the cookie 
func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}




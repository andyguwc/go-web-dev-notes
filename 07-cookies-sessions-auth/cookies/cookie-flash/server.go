
/* Using cookie to power flash messages (only once for example upon form submission)


*/

package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// Base64 URL encoding of the message

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

// get the cookie value and then delete it (by creating one with same name and MaxAge -1)
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}

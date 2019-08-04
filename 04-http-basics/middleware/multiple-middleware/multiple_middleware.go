/* 
Multiple Middlewares

In the first middleware, check whether the content type is JSON. If not, don't
allow the request to proceed

In the second middleware, add a timestamp called Server-Time (UTC) to the
response cookie

Great example for JSON code too
*/


package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"log"
)

type city struct {
	Name string
	Area uint64
}

// middleware to check content type as JSON
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		// For checking, we are using r.Header.GET (content type)
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			
			return 
		}
		handler.ServeHTTP(w,r)
	})
}

// middleware to add server timestamp for response cookie 
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w,r)

		// setting cookie to each and every response
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(),10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}


// mainLogic is just a simple POST function 
func mainLogic(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method == "POST" {
		var tempCity city 
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with area of %d sq miles!\n",tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}

}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLogicHandler)))
	http.ListenAndServe(":8000", nil)
}

// curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}'

// if we remove the content-type it won't work 
// curl -i -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}'



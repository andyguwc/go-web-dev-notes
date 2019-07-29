/*
Form property on the Request object will contain values from the URL query along with values submitted as a POST or PUT body

*/
package form 

import (
	"fmt"
	"net/http"
)
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	name := r.FormValue("name")
	fmt.Println(name)

}

// accessing multiple values from a form field
// for example check boxes 

func anotherHandler(w http.ResponseWriter, r *http.Request) {
	var maxMemory int64 = 16 << 20
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
	fmt.Println(err)
	}
	for _, v := range r.PostForm["names"] {
	fmt.Println(v)
	}
}

// Passing value https://github.com/GoesToEleven/golang-web-dev/tree/master/027_passing-data
// form methods post vs. get 
// get go through url 
// post go through body 



/* Refactor to use MVC structure

Controller has a session to Mongo

Make handlers a method on the usercontroller
*/


package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/andyguwc/go-course/golang-web-dev/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}


// return mgo.Session to pass in the UserController
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

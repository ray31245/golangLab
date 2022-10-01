package main

import (
	"Bevis/src/webDev/042_mongodb/05_mongodb/01_update-user-controller/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	// Get a UserController instanse
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.CreateUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// check if connection error, is mongo runing?
	if err != nil {
		panic(err)
	}
	return s
}

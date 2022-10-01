package main

import (
	"Bevis/src/webDev/042_mongodb/05_mongodb/05_update-user-controllers-delete/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	session := getSession()
	defer session.Close()
	uc := controllers.NewUserController(session)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://my_mongo:27017/")
	if err != nil {
		panic(err)
	}
	return session
}

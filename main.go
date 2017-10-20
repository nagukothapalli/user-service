package main

import (
	"net/http"
	"user-service/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {

	//User-Service API
	userRestController := controllers.NewUserRestController()
	//Building the Router
	router := httprouter.New()
	router.GET("/", userRestController.Index)
	router.GET("/user/:id", userRestController.GetUserById)
	router.GET("/users/all", userRestController.GetAllUsers)
	router.POST("/create", userRestController.CreateUser)
	http.ListenAndServe(":8080", router)

}

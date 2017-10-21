package controllers

import (
	//standard lib packages
	"encoding/json"
	"fmt"
	"net/http"
	//custome
	"user-service/models"
	"user-service/services"
	//extenal libs
	"github.com/julienschmidt/httprouter"
)

var userService *services.UserService

func init() {
	userService = services.NewuserService()
}

type UserRestController struct {
}

//TO get new controller
func NewUserRestController() *UserRestController {
	return &UserRestController{}
}

//Index Page
func (userRestController UserRestController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(`<h1> Hello , User-Service</h1>`))

}

//Get user details by ID
func (userRestController UserRestController) GetUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	userResponse := userService.GetUser(id)
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(&userResponse)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(userResponse.StatusCode)
	fmt.Fprintf(w, "%s", uj)
}

//curl -vvv -H "Content-Type: application/json" -X POST -d'{"Name":"Nagu Kothapalli","Gender":"male","Age":50,"Id":"a@b.com"}' http://localhost:8080/create
func (userRestController UserRestController) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	newuser := models.User{}
	json.NewDecoder(r.Body).Decode(&newuser)
	userResponse := userService.CreateUser(newuser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(userResponse.StatusCode)

	uj, _ := json.Marshal(&userResponse)
	fmt.Fprintf(w, "%s", uj)

}

func (userRestController UserRestController) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userResponse := userService.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(userResponse.StatusCode)
	uj, _ := json.Marshal(&userResponse)
	w.Write([]byte(uj))

}

func (userRestController UserRestController) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

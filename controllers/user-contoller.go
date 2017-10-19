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
	user := userService.GetUser(id)
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(user)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", uj)
}

//curl -vvv -H "Content-Type: application/json" -X POST -d'{"Name":"Nagu Kothapalli","Gender":"male","Age":50,"Id":"a@b.com"}' http://localhost:8080/create
func (userRestController UserRestController) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	newuser := models.User{}
	json.NewDecoder(r.Body).Decode(&newuser)
	fmt.Println(newuser)
	userService.CreateUser(newuser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	uj, _ := json.Marshal(&newuser)
	fmt.Fprintf(w, "%s", uj)

}

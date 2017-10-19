package services

import (
	"go-learning/01_helloworld/model"
	"user-service/backend"
	"user-service/models"
)

type UserService struct {
}

func NewuserService() *UserService {
	return &UserService{}
}

func (userService UserService) CreateUser(user models.User) {

	con := backend.GetNewDBConnection()
	err := con.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	stmtIns, err := con.Prepare("INSERT INTO User VALUES( ?, ?,?,? )") // ? = placeholder
	_, err = stmtIns.Exec(user.Id, user.Name, user.Gender, user.Age)
	defer stmtIns.Close()
	defer con.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

func (userService UserService) GetUser(Id string) models.User {

	var user models.User
	con := backend.GetNewDBConnection()

	err := con.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	stmtOut, err := con.Prepare("Select Id,Name,Gender,Age FROM User where id = ?") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}

	defer stmtOut.Close()
	defer con.Close()

	err = stmtOut.QueryRow(Id).Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return user
}

func (userService UserService) GetAllUser() []model.User {
	return nil
}

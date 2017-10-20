package services

import (
	"database/sql"
	"user-service/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {
}

func NewuserService() *UserService {
	return &UserService{}
}

func (userService UserService) CreateUser(user models.User) {

	con := getNewDBConnection()
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
	con := getNewDBConnection()

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

func (userService UserService) GetAllUser() []models.User {
	return nil

}

func getNewDBConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:nagu123@/gotest")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db

}

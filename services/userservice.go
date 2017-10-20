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

func (userService UserService) GetAllUsers() []models.User {

	users := make([]models.User, 0)

	con := getNewDBConnection()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	results, err := con.Query("Select Id,Name,Gender,Age FROM User") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app

	}

	for results.Next() {

		var user models.User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		users = append(users, user)
	}
	return users

}

// for getting the DB object
func getNewDBConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:nagu123@/gotest")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

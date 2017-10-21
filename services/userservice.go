package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"user-service/models"

	_ "github.com/go-sql-driver/mysql"
)

type userResponse models.UserRestResponse

type UserService struct {
}

func NewuserService() *UserService {
	return &UserService{}
}

func (userService UserService) CreateUser(user models.User) userResponse {

	con := getNewDBConnection()
	err := con.Ping()
	// Check DB  connection is fine or not
	if err != nil {
		return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	}
	isExists, err := isUserExists(user.Id)

	if err != nil {
		return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	}
	if isExists {
		return buildUserResponse(false, http.StatusConflict, "User Already exist with same id", nil)
	}
	stmtIns, _ := con.Prepare("INSERT INTO User VALUES( ?, ?,?,? )") // ? = placeholder
	_, err = stmtIns.Exec(user.Id, user.Name, user.Gender, user.Age)

	var response userResponse
	if err != nil {
		response = buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	} else {
		response = buildUserResponse(false, http.StatusCreated, "User Created Successfully", user)
	}
	defer stmtIns.Close()
	defer con.Close()
	return response
}

func (userService UserService) GetUser(Id string) userResponse {

	var user models.User
	con := getNewDBConnection()
	err := con.Ping()
	// Check DB  connection is fine or not
	if err != nil {
		return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	}
	isExists, err := isUserExists(Id)

	if !isExists {
		return buildUserResponse(true, http.StatusNotFound, "User not created yet with passed id", nil)
	}
	if err != nil {
		return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	}

	stmtOut, _ := con.Prepare("Select Id,Name,Gender,Age FROM User where id = ?") // ? = placeholder

	err = stmtOut.QueryRow(Id).Scan(&user.Id, &user.Name, &user.Gender, &user.Age)
	var response userResponse

	if err != nil {
		response = buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	} else {
		response = buildUserResponse(false, http.StatusOK, "User returned Succesfully", user)
	}

	defer stmtOut.Close()
	defer con.Close()
	return response
}

func (userService UserService) GetAllUsers() userResponse {

	users := make([]models.User, 0)

	con := getNewDBConnection()
	err := con.Ping()
	// Check DB  connection is fine or not
	if err != nil {
		return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	}

	results, err := con.Query("Select Id,Name,Gender,Age FROM User") // ? = placeholder
	if err != nil {
		return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
	}

	for results.Next() {

		var user models.User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.Id, &user.Name, &user.Gender, &user.Age)

		if err != nil {
			return buildUserResponse(true, http.StatusInternalServerError, err.Error(), nil)
		}
		// and then print out the tag's Name attribute
		users = append(users, user)
	}
	if len(users) == 0 {
		return buildUserResponse(false, http.StatusOK, "Users not created yet.", users)
	}
	return buildUserResponse(false, http.StatusOK, "All users returning successfully", users)

}

// helper Methods
func getNewDBConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:nagu123@/gotest")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

func buildUserResponse(isError bool, statusCode int, message string, data interface{}) userResponse {
	resp := userResponse{isError, statusCode, message, data}
	return resp
}

func isUserExists(Id string) (bool, error) {

	var count int
	con := getNewDBConnection()
	query := fmt.Sprintf("SELECT COUNT(*) as count FROM  User where Id=%s", Id)
	rows, err := con.Query(query)

	for rows.Next() {
		rows.Scan(&count)
	}
	defer con.Close()
	return count != 0, err
}

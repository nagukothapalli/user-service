package models

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

type UserRestResponse struct {
	Error      bool        `json:"error"`
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"response"`
}

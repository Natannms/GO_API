package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"Nome"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users []User

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Natannms/GO_API/database"
	"github.com/Natannms/GO_API/models"
	"github.com/gorilla/mux"
)

func Home(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("home")
}

func AllUsers(r http.ResponseWriter, rq *http.Request) {
	var users []models.User
	database.Conn().Find(&users)
	json.NewEncoder(r).Encode(users)
}
func GetUser(r http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id := vars["id"]
	var user models.User
	database.Conn().First(&user, id)

	json.NewEncoder(r).Encode(user)

}
func CreateUser(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("Usuario Criado com Sucesso")
}

func UpdateUser(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("Usuario Atualizado com Sucesso")
}

func DeleteUser(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("Usuario Deletado com Sucesso")
}

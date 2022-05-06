package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Natannms/GO_API/models"
)

func Home(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("home")
}

func AllUsers(r http.ResponseWriter, rq *http.Request) {
	//pesquisar no banco de dados todos os usuarios
	ListaDeUsuarios := []models.User{
		models.User{ID: 1, Name: "Natan", Email: "lara@gmail.com", Password: "123"},
		models.User{ID: 2, Name: "Lara", Email: "natan@gmail.com", Password: "123"},
	}

	models.Users = ListaDeUsuarios
	json.NewEncoder(r).Encode(models.Users)
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

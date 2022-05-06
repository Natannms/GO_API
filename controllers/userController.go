package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Natannms/GO_API/models"
	"github.com/gorilla/mux"
)

func Home(r http.ResponseWriter, rq *http.Request) {
	fmt.Println("home")
}

func AllUsers(r http.ResponseWriter, rq *http.Request) {

	json.NewEncoder(r).Encode(models.Users)
}
func GetUser(r http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	id := vars["id"] // Pega o parametro id da url

	for _, user := range models.Users {
		//percorre todo o array
		if strconv.Itoa(user.Id) == id {
			//compara id para retornar as informações
			json.NewEncoder(r).Encode(user)
		}
	}
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

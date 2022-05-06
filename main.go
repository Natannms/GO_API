package main

import (
	"fmt"

	"github.com/Natannms/GO_API/models"
	"github.com/Natannms/GO_API/routes"
)

func main() {
	//pesquisar no banco de dados todos os usuarios
	ListaDeUsuarios := []models.User{
		models.User{Id: 1, Name: "Natan", Email: "lara@gmail.com", Password: "123"},
		models.User{Id: 2, Name: "Lara", Email: "natan@gmail.com", Password: "123"},
	}

	models.Users = ListaDeUsuarios

	fmt.Println("Running in port http://localhost:8000") // Mensagem com link para executar no nvegador
	routes.HandleRequest()
}

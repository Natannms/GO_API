package routes

import (
	"log"
	"net/http"

	"github.com/Natannms/GO_API/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home) // set router Home
	http.HandleFunc("/user", controllers.AllUsers)
	http.HandleFunc("/user", controllers.CreateUser)
	http.HandleFunc("/user/:1", controllers.UpdateUser)
	http.HandleFunc("/user/:1", controllers.DeleteUser)
	log.Fatal(http.ListenAndServe(":8000", nil)) //Escutando a porta 8000
}

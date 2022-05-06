package routes

import (
	"log"
	"net/http"

	"github.com/Natannms/GO_API/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/user", controllers.AllUsers)
	r.HandleFunc("/user", controllers.CreateUser)
	r.HandleFunc("/user/{id}", controllers.GetUser)
	r.HandleFunc("/user/{id}", controllers.UpdateUser)
	r.HandleFunc("/user/{id}", controllers.DeleteUser)

	log.Fatal(http.ListenAndServe(":8000", r))
}

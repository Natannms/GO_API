package routes

import (
	"log"
	"net/http"

	"github.com/Natannms/GO_API/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	r := mux.NewRouter()
	// r.Use(middleware.HeaderMiddleare)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/user", controllers.AllUsers).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

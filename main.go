package main

import (
	"fmt"

	"github.com/Natannms/GO_API/routes"
)

func main() {

	fmt.Println("Running in port http://localhost:8000") // Mensagem com link para executar no nvegador
	routes.HandleRequest()
}

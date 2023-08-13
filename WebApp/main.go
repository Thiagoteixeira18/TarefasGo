package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	utils.CarregarTemplate()
	r := router.Gerar()

	fmt.Println("Escutando na porta 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
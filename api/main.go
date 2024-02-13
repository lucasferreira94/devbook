package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucasferreira94/devbook/tree/main/api/src/router"
)

func main() {
	fmt.Println("rodando api")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}

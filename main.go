package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarConfig()
	fmt.Print(config.Porta)
	fmt.Print(config.StringConexaoBanco)
	fmt.Println("servidor rodando")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":50000", r))

}

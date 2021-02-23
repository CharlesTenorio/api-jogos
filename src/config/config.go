package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

func CarregarConfig() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Porta, error = strconv.Atoi(os.Getenv("API_PORTA"))
	if error != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"),
		os.Getenv("PGUSER"), os.Getenv("PASSWORD_DB"), os.Getenv("PGDATABASE"))

}

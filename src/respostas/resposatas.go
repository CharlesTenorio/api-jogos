package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)

	}

}

func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json: "Error"`
	}{
		Error: erro.Error(),
	})
}

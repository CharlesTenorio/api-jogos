package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ListaTodosTimes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lista todos os times"))
}

func ListaTtime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("busca time"))

}

func CriarTime(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var time models.Time
	if erro = json.Unmarshal(corpoRequest, &time); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repository.NovoRepoTimes(db)

	timeID, erro := repositorio.Criar(time)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", timeID)))

}

func AutalizarTimes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizar os times"))
}

func Deletartime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar busca time"))

}

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

func ListaTodasLigas(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("lista todas as ligas"))
}

func ListaLiga(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("busca LIGA"))

}

func CriarLiga(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}
	var liga models.Liga
	if erro = json.Unmarshal(corpoRequest, &liga); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repository.NovoRepoLiga(db)

	ligaID, erro := repositorio.CriarLiga(liga)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", ligaID)))

}

func AutalizarLiga(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizar os times"))
}

func DeletarLiga(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar busca time"))

}

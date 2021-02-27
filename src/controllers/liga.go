package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
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
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return

	}
	var liga models.Liga
	if erro = json.Unmarshal(corpoRequest, &liga); erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepoLiga(db)

	liga.Id, erro = repositorio.CriarLiga(liga)
	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, liga)

}

func AutalizarLiga(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizar os times"))
}

func DeletarLiga(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar busca time"))

}

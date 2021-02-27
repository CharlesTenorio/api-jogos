package repository

import (
	"api/src/models"
	"database/sql"
)

type liga struct {
	db *sql.DB
}

func NovoRepoLiga(db *sql.DB) *liga {
	return &liga{db}

}

func (repoliga liga) CriarLiga(liga models.Liga) (uint64, error) {

	statment, erro := repoliga.db.Prepare("INSERT INTO liga (id, nome, cc, has_toplist, has_leaguetable) VALUES(?, ?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statment.Close()
	resultado, erro := statment.Exec(liga.Id, liga.Nome, liga.Cc, liga.Has_toplist, liga.Has_leaguetable)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

package repository

import (
	"api/src/models"
	"database/sql"
)

type times struct {
	db *sql.DB
}

func NovoRepoTimes(db *sql.DB) *times {
	return &times{db}

}

func (repotime times) Criar(time models.Time) (uint64, error) {
	statment, erro := repotime.db.Prepare("insert into time (id_time, nome, cc, imagem_id ) valeus(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()
	resultado, erro := statment.Exec(time.IdTime, time.Nome, time.Cc, time.ImageId)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

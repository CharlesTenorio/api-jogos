package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" //drive possgresql
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("postgres", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
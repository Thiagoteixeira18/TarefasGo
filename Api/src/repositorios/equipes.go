package repositorios

import (
	"api/src/equipe"
	"database/sql"
)

type Equipe struct {
	db *sql.DB
}

func NovoRepositorioDeEquipes(db *sql.DB) *Equipe {
	return &Equipe{db}
}

func (repositorio Equipe) CriarEquipe(equipe equipe.Equipes) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into equipes (nome, descricao, autor_id) value(?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(equipe.Nome, equipe.Descricao, equipe.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

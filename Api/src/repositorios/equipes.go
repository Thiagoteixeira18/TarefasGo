package repositorios

import (
	"api/src/equipe"
	"database/sql"
	"fmt"
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

func (repositorio Equipe) Buscar(nomeDaEquipe string) ([]equipe.Equipes, error) {
	nomeDaEquipe = fmt.Sprintf("%%%s%%", nomeDaEquipe)

	linhas, erro := repositorio.db.Query(
		"select id, nome, descricao, autor_id from equipes where nome LIKE ?",
		nomeDaEquipe,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var equipes []equipe.Equipes
	for linhas.Next() {
		var equipe equipe.Equipes

		if erro = linhas.Scan(
			&equipe.Id,
			&equipe.Nome,
			&equipe.Descricao,
			&equipe.AutorId,
		); erro != nil {
			return nil, erro
		}
		equipes = append(equipes, equipe)
	}

	return equipes, nil
}

func (repositorio Equipe) BuscarPorId(equipeId uint64) (equipe.Equipes, error) {
	linha, erro := repositorio.db.Query("select id, nome, descricao, autor_id from equipes where id = ?", equipeId)
	if erro != nil {
		return equipe.Equipes{}, erro
	}
	defer linha.Close()

	var Equipe equipe.Equipes

	if linha.Next() {
		if erro = linha.Scan(
			&Equipe.Id,
			&Equipe.Nome,
			&Equipe.Descricao,
			&Equipe.AutorId,
		); erro != nil {
			return equipe.Equipes{}, erro
		}
	}

	return Equipe, nil
}

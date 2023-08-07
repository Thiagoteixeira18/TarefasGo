package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Equipe struct {
	db *sql.DB
}

func NovoRepositorioDeEquipes(db *sql.DB) *Equipe {
	return &Equipe{db}
}

func (repositorio Equipe) CriarEquipe(equipe modelos.Equipes) (uint64, error) {
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

func (repositorio Equipe) Buscar(nomeDaEquipe string) ([]modelos.Equipes, error) {
	nomeDaEquipe = fmt.Sprintf("%%%s%%", nomeDaEquipe)

	linhas, erro := repositorio.db.Query(
		"select id, nome, descricao, autor_id from equipes where nome LIKE ?",
		nomeDaEquipe,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var equipes []modelos.Equipes
	for linhas.Next() {
		var equipe modelos.Equipes

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

func (repositorio Equipe) BuscarPorId(equipeId uint64) (modelos.Equipes, error) {
	linha, erro := repositorio.db.Query("select id, nome, descricao, autor_id from equipes where id = ?", equipeId)
	if erro != nil {
		return modelos.Equipes{}, erro
	}
	defer linha.Close()

	var Equipe modelos.Equipes

	if linha.Next() {
		if erro = linha.Scan(
			&Equipe.Id,
			&Equipe.Nome,
			&Equipe.Descricao,
			&Equipe.AutorId,
		); erro != nil {
			return modelos.Equipes{}, erro
		}
	}

	return Equipe, nil
}

func (repositorio Equipe) AtualizarEquipe(equipeId uint64, Equipe modelos.Equipes) error {
	statement, erro := repositorio.db.Prepare("update equipes set nome = ?, descricao = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Equipe.Nome, Equipe.Descricao, equipeId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) DeletarEquipe(equipeId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from equipes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(equipeId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Equipe) CriarTarefaDeEquipe(tarefa modelos.Tarefas, equipeId uint64, usuarioId uint64) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into tarefas_equipe (tarefa, observacao, prazo, autor_id, equipes_id) value(?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tarefa.Tarefa, tarefa.Obsevacao, tarefa.Prazo, tarefa.AutorId, equipeId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Equipe) BuscarTarefasDaEquipe(equipeId uint64) ([]modelos.Tarefas, error) {
	linhas, erro := repositorio.db.Query("SELECT tarefa, observacao, prazo FROM tarefas_equipe WHERE equipes_id = ?", equipeId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tarefasDaEquipe []modelos.Tarefas

	for linhas.Next() {
		var tarefa modelos.Tarefas

		if erro = linhas.Scan(
			&tarefa.Tarefa,
			&tarefa.Obsevacao,
			&tarefa.Prazo,
		); erro != nil {
			return nil, erro
		}

		tarefasDaEquipe = append(tarefasDaEquipe, tarefa)
	}

	return tarefasDaEquipe, nil
}

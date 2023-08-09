package repositorios

import (
    "api/src/modelos"
    "database/sql"
    "errors"
)

type Tarefas struct {
    db *sql.DB
}

func NovoRepositorioDeTarefas(db *sql.DB) *Tarefas {
    return &Tarefas{db}
}

func (repositorio Tarefas) CriarTarefa(tarefa modelos.Tarefas) (uint64, error) {
    statement, erro := repositorio.db.Prepare(
        "insert into tarefas (tarefa, observacao, prazo, autor_id) value(?, ?, ?, ?)",
    )
    if erro != nil {
        return 0, erro
    }
    defer statement.Close()

    resultado, erro := statement.Exec(tarefa.Tarefa, tarefa.Obsevacao, tarefa.Prazo, tarefa.AutorId)
    if erro != nil {
        return 0, erro
    }

    ultimoIdInserido, erro := resultado.LastInsertId()
    if erro != nil {
        return 0, erro
    }

    return uint64(ultimoIdInserido), nil
}

func (repositorio Tarefas) BuscarTarefa(tarefaId uint64) (modelos.Tarefas, error) {
    linha, erro := repositorio.db.Query("select id, autor_id, tarefa, observacao, prazo from tarefas where id = ?",
        tarefaId,
    )
    if erro != nil {
        return modelos.Tarefas{}, erro
    }
    defer linha.Close()

    var tarefa modelos.Tarefas

    if linha.Next() {
        if erro = linha.Scan(
            &tarefa.Id,
            &tarefa.AutorId,
            &tarefa.Tarefa,
            &tarefa.Obsevacao,
            &tarefa.Prazo,          
        ); erro != nil {
            return modelos.Tarefas{}, erro
        }
    }

    return tarefa, nil

}

func (repositorio Tarefas) BuscarPorUsuario(usuarioId uint64) ([]modelos.Tarefas, error) {
    linhas, erro := repositorio.db.Query(`
    select t.*, u.nick from tarefas t
    join usuarios u on u.id = t.autor_id
    where t.autor_id = ?`,
        usuarioId,
    )
    if erro != nil {
        return nil, erro
    }
    defer linhas.Close()

    var Tarefa []modelos.Tarefas

    for linhas.Next() {
        var tarefa modelos.Tarefas

        if erro = linhas.Scan(
            &tarefa.Id,
            &tarefa.Tarefa,
            &tarefa.Obsevacao,
            &tarefa.AutorId,
            &tarefa.Prazo,
            &tarefa.AutorNick,          
        ); erro != nil {
            return nil, erro
        }
        if usuarioId != tarefa.AutorId {
            return nil, errors.New("Não é possível buscar tarefas de outro usuario!")
        }

        Tarefa = append(Tarefa, tarefa)
    }
    return Tarefa, nil
}

func (repositorio Tarefas) Deletar(tarefaId uint64) error {
    statement, erro := repositorio.db.Prepare("delete from tarefas where id = ? ")
    if erro != nil {
        return erro
    }
    defer statement.Close()

    if _, erro := statement.Exec(tarefaId); erro != nil {
        return erro
    }

    return nil

}

func (repositorio Tarefas) Atualizar(tarefaId uint64, tarefaAtualziada modelos.Tarefas) error {
    statement, erro := repositorio.db.Prepare("update tarefas set tarefa = ?, observacao = ?, prazo = ? where id = ?")
    if erro != nil {
        return erro
    }
    defer statement.Close()

    if _, erro := statement.Exec(tarefaAtualziada.Tarefa, tarefaAtualziada.Obsevacao, tarefaAtualziada.Prazo, tarefaId); erro != nil {
        return erro
    }

    return nil
}


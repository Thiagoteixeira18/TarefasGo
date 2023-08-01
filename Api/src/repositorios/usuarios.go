package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeusuairos cria um novo repositorio de usuarios
func NovoRepositorioDeusuairos(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuarios) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Usuarios) BuscarPorId(usuarioId uint64) (modelos.Usuarios, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email from usuarios where id = ?",
		usuarioId,
	)
	if erro != nil {
		return modelos.Usuarios{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuarios

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); erro != nil {
			return modelos.Usuarios{}, erro
		}
	}

	return usuario, nil

}

func (repositorio Usuarios) Atualizar(Id uint64, usuario modelos.Usuarios) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, Id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(usuarioId uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuarios, error) {
	linhas, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)
	if erro != nil {
		return modelos.Usuarios{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuarios

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.Id, &usuario.Senha); erro != nil {
			return modelos.Usuarios{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuarios

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil

}

func (repositorio Usuarios) AtualizarSenha(usuarioId uint64, senhaComHash string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(senhaComHash, usuarioId); erro != nil {
		return erro
	}

	return nil
}

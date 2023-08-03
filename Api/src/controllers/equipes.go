package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/equipe"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarEquipes(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var equipe equipe.Equipes
	if erro = json.Unmarshal(corpoRequest, &equipe); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	equipe.AutorId = usuarioId

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	equipe.Id, erro = repositorio.CriarEquipe(equipe)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusCreated, equipe)

}

func BuscarEquipes(w http.ResponseWriter, r *http.Request) {}

func BuscarEquipe(w http.ResponseWriter, r *http.Request) {}

func AtualizarEquipe(w http.ResponseWriter, r *http.Request) {}

func DeletarEquipe(w http.ResponseWriter, r *http.Request) {}

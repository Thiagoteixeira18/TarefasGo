package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CriarTarefa(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var tarefa modelos.Tarefas
	if erro = json.Unmarshal(corpoRequest, &tarefa); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	tarefa.AutorId = usuarioId

	if erro = tarefa.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	layout := "02-01-2006"
	tarefaPrazo, erro := time.Parse(layout, tarefa.Prazo)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	tarefa.Prazo = tarefaPrazo.Format(time.RFC3339)

	diaDaSemana := tarefaPrazo.Weekday().String()
	tarefa.Prazo = fmt.Sprintf("%s (%s)", tarefaPrazo.Format("02/01/2006"), diaDaSemana)

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeTarefas(db)
	tarefa.Id, erro = repositorio.CriarTarefa(tarefa)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusCreated, tarefa)

}

func BuscarTarefas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("buscando")
}

func BuscarTarefa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaId, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeTarefas(db)
	tarefa, erro := repositorio.BuscarTarefa(tarefaId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, tarefa)

}

func EditarTarefa(w http.ResponseWriter, r *http.Request) {

}

func DeletarTarefa(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametros := mux.Vars(r)
	tarefaId, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeTarefas(db)
	tarefaSalvaNoBanco, erro := repositorio.BuscarTarefa(tarefaId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioId != tarefaSalvaNoBanco.AutorId {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("Não pode apagar uma tarefa que não é sua!"))
		return
	}

	if erro = repositorio.Deletar(tarefaId); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

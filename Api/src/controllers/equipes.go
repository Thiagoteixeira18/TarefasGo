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
	"strings"
	"time"

	"github.com/gorilla/mux"
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

	var equipe modelos.Equipes
	if erro = json.Unmarshal(corpoRequest, &equipe); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	equipe.AutorId = usuarioId

	if erro = equipe.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

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

func BuscarEquipes(w http.ResponseWriter, r *http.Request) {
	nomeDaEquipe := strings.ToLower(r.URL.Query().Get("equipes"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	equipes, erro := repositorio.Buscar(nomeDaEquipe)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, equipes)
}

func BuscarEquipe(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	equipe, erro := repositorio.BuscarPorId(equipeId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, equipe)
}

func AtualizarEquipe(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametros := mux.Vars(r)
	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	equipeSalvaNoBanco, erro := repositorio.BuscarPorId(equipeId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioId != equipeSalvaNoBanco.AutorId {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("Não é possível atualizar uma equipe que você não seja o administrador!"))
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var Equipe modelos.Equipes

	if erro = json.Unmarshal(corpoRequest, &Equipe); erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	if erro = Equipe.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarEquipe(equipeId, Equipe); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}

func DeletarEquipe(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	parametros := mux.Vars(r)
	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	equipeSalvaNoBanco, erro := repositorio.BuscarPorId(equipeId)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioId != equipeSalvaNoBanco.AutorId {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("Não é possível deletar uma equipe que voçê não seja adiministrador"))
		return
	}

	if erro = repositorio.DeletarEquipe(equipeId); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}

func CriarTarefaDeEquipe(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var tarefa modelos.Tarefas

	if erro = json.Unmarshal(corpoRequest, &tarefa); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	tarefa.AutorId = usuarioId

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
	defer db.Close()

	parametros := mux.Vars(r)
	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	repositorioEquipe := repositorios.NovoRepositorioDeEquipes(db)
	tarefa.Id, erro = repositorioEquipe.CriarTarefaDeEquipe(tarefa, equipeId, usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, tarefa)

}

func BuscarTarefasDaEquipe(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	tarefas, erro := repositorio.BuscarTarefasDaEquipe(equipeId)
	if erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, tarefas)
}

func BuscarTarefaDaEquipe(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	tarefaId, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	tarefa, erro := repositorio.BuscarTarefaDaEquipe(tarefaId, equipeId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return 
	}
	
	respostas.JSON(w, http.StatusOK, tarefa)
}

func EditarTarefaDaEquipe(w http.ResponseWriter, r *http.Request) {
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

	equipeId, erro := strconv.ParseUint(parametros["equipeId"], 10, 64)
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

	repositorio := repositorios.NovoRepositorioDeEquipes(db)
	TarefaSalvaNoBanco, erro := repositorio.BuscarTarefaDaEquipe(tarefaId, equipeId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if TarefaSalvaNoBanco.AutorId != usuarioId {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("Não é possivel atualizar uma tarefa que não tenha sido você quem criou"))
		return 
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusForbidden, erro)
		return
	}

	var Tarefa modelos.Tarefas
	if erro = json.Unmarshal(corpoRequest, &Tarefa); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.EditarTarefaDaEquipe(equipeId, tarefaId, Tarefa); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}

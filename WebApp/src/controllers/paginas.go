package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.ExecutarTemplete(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/tarefas", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var tarefas []modelos.Tarefas
	if erro = json.NewDecoder(response.Body).Decode(&tarefas); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "home.html", tarefas)

}

func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioId, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "perfil.html", usuario)
}

func CarregarPaginaDeEdicaoDoUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan modelos.Usuario)
	go modelos.BuscaDadosUsuario(canal, usuarioId, r)
	usuario := <-canal

	if usuario.Id == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao buscar usuário"})
		return
	}

	utils.ExecutarTemplete(w, "editar-usuario.html", usuario)

}

func CarregarPaginaDeEdicaoDeTarefa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaId, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/tarefas/%d", config.APIURL, tarefaId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var tarefa modelos.Tarefas

	if erro = json.NewDecoder(response.Body).Decode(&tarefa); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "editar-tarefa.html", tarefa)
}

func CarregarPaginaDeEdicaoDoSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplete(w, "editar-senha.html", nil)
}

func CarregarPaginaDeEquipes(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan []modelos.Equipes)
	go modelos.BuscaEquipesDoUsuario(canal, usuarioId, r)
	equipesCarregadas := <-canal

	var equipes []modelos.Equipes
	if equipesCarregadas == nil {
		equipes = []modelos.Equipes{}
	} else {
		equipes = equipesCarregadas
	}

	utils.ExecutarTemplete(w, "equipes.html", equipes)
}

func CarregarPaginaDeEdicaoDeEquipe(w http.ResponseWriter, r *http.Request) {
	paramentros := mux.Vars(r)
	equipeId, erro := strconv.ParseUint(paramentros["equipeId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/equipes/%d", config.APIURL, equipeId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var equipe modelos.Equipes

	if erro = json.NewDecoder(response.Body).Decode(&equipe); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "editar-equipe.html", equipe)
}

func CarregarPaginaDeEdicaoDeTarefaDeEquipe(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaId, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/tarefas/%d/equipes", config.APIURL, tarefaId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var tarefa modelos.Tarefas

	if erro = json.NewDecoder(response.Body).Decode(&tarefa); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplete(w, "editar-tarefa-equipe.html", tarefa)
}

package rotas

import (
	"api/src/controllers"
	"net/http"
)

var RotaDeEquipes = []Rota{
	{
		Uri:                   "/equipes",
		Metodo:                http.MethodPost,
		Funcao:                controllers.CriarEquipes,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarEquipes,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarEquipe,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}",
		Metodo:                http.MethodPut,
		Funcao:                controllers.AtualizarEquipe,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}",
		Metodo:                http.MethodDelete,
		Funcao:                controllers.DeletarEquipe,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}/tarefas",
		Metodo:                http.MethodPost,
		Funcao:                controllers.CriarTarefaDeEquipe,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}/tarefas",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarTarefasDaEquipe,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}/tarefas/{tarefaId}",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarTarefaDaEquipe,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/equipes/{equipeId}/tarefas/{tarefaId}",
		Metodo:                http.MethodPut,
		Funcao:                controllers.EditarTarefaDaEquipe,
		RequerAutententicacao: true,
	},
	{
        Uri:                   "/equipes/{equipeId}/tarefas/{tarefaId}",
        Metodo:                http.MethodDelete,
        Funcao:                controllers.DeletarTarefaDaEquipe,
        RequerAutententicacao: true,
    },
	{
        Uri:                   "/equipes/{equipeId}/adicionar/{usuarioId}",
        Metodo:                http.MethodPost,
        Funcao:                controllers.AdicionarUsuario,
        RequerAutententicacao: true,
    },
	{
        Uri:                   "/equipes/{equipeId}/remover/{usuarioId}",
        Metodo:                http.MethodDelete,
        Funcao:                controllers.RemoverUsuario,
        RequerAutententicacao: true,
    },
	{
        Uri:                   "/equipes/{equipeId}/usuario/{usuarioId}",
        Metodo:                http.MethodGet,
        Funcao:                controllers.BuscarUsuarioDaEquipe,
        RequerAutententicacao: true,
    },
	{
        Uri:                   "/equipes/{equipeId}/usuarios",
        Metodo:                http.MethodGet,
        Funcao:                controllers.BuscarUsuariosDaEquipe,
        RequerAutententicacao: true,
    },
}
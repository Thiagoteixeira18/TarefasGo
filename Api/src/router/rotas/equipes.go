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
}

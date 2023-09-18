package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotaEquipes = []Rota{
	{
		Uri:                "/equipe",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEquipes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipe",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEquipes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}/editar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/tarefas/{tarefaId}/equipes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeTarefaDeEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/tarefas/{tarefaId}/equipe",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarTarefaDeEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscainformacoesDaEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}/tarefas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.NovaTarefasDeEquipe,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/equipes/{equipeId}/tarefas/{tarefaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ConcluirEDeletarTarefaDeEquipe,
		RequerAutenticacao: true,
	},
}

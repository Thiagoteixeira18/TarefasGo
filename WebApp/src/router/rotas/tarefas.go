package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasTarefas = []Rota{
	{
		Uri:                "/tarefas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTarefa,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/tarefas/{tarefaId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ConcluirTarefa,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/tarefas/{tarefaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarTarefa,
		RequerAutenticacao: true,
	},
}

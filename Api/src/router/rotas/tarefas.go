package rotas

import (
	"api/src/controllers"
	"net/http"
)

var RotasTarefas = []Rota{
	{
		Uri:                   "/tarefas",
		Metodo:                http.MethodPost,
		Funcao:                controllers.CriarTarefa,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/tarefas",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarTarefas,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/tarefas/{tarefaId}",
		Metodo:                http.MethodPut,
		Funcao:                controllers.EditarTarefa,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/tarefas/{tarefaId}",
		Metodo:                http.MethodDelete,
		Funcao:                controllers.DeletarTarefa,
		RequerAutententicacao: true,
	},
}

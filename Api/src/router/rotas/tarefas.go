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
		Uri:                   "/tarefas/{tarefaId}",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarTarefa,
		RequerAutententicacao: true,
	},
	{
		Uri:                   "/usuarios/{usuarioId}/tarefas",
		Metodo:                http.MethodGet,
		Funcao:                controllers.BuscarTarefasDoUsuario,
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

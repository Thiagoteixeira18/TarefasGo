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
}

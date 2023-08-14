package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotaUsuarios = []Rota{
	{
		Uri:    "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
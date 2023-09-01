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
	{
		Uri:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
}
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
		Uri:                "/equipes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarEquipes,
		RequerAutenticacao: true,
	},
}
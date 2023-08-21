package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotaPaginaPrincipal = Rota{
	Uri:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
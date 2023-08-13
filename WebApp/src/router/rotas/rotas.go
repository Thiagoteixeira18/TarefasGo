package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}


func Configurar(router *mux.Router) *mux.Router {
	rotas := RotasLogin

	for _, rota := range rotas {
		router.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	return router
}
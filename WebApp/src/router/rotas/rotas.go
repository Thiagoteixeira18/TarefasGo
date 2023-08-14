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
	rotas = append(rotas, rotaUsuarios...)

	for _, rota := range rotas {
		router.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

type Usuario struct {
	Id      uint64    `json:"id"`
	Nome    string    `json:"nome"`
	Nick    string    `json:"nick"`
	Email   string    `json:"email"`
	Tarefas []Tarefas `json:"tarefas"`
	Equipes []Equipes `json:"equipes"`
}

func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {
    canalUsuario := make(chan Usuario)
    canalTarefas := make(chan []Tarefas)
    canalEquipes := make(chan []Equipes)

    go BuscaDadosUsuario(canalUsuario, usuarioId, r)
    go BuscaTarefasDoUsuaro(canalTarefas, usuarioId, r)
    go BuscaEquipesDoUsuario(canalEquipes, usuarioId, r)

    var (
        usuarioCompleto Usuario
        tarefas         []Tarefas
        equipes         []Equipes
        usuarioErro     error
    )

    for i := 0; i < 3; i++ {
        select {
        case usuarioCarregado := <-canalUsuario:
            if usuarioCarregado.Id == 0 {
                usuarioErro = errors.New("Erro ao buscar usuário")
            }
            usuarioCompleto = usuarioCarregado

        case tarefasCarregadas := <-canalTarefas:
            if tarefasCarregadas == nil {
				tarefas = []Tarefas{}
			} else {
                return Usuario{}, errors.New("Erro ao buscar tarefas")
            }
            tarefas = tarefasCarregadas

        case equipesCarregadas := <-canalEquipes:
            if equipesCarregadas == nil {
                equipes = []Equipes{}
            } else {
                equipes = equipesCarregadas
            }
        }
    }

    if usuarioErro != nil {
        return Usuario{}, usuarioErro
    }

    usuarioCompleto.Tarefas = tarefas
    usuarioCompleto.Equipes = equipes

    return usuarioCompleto, nil
}


func BuscaDadosUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

func BuscaTarefasDoUsuaro(canal chan<- []Tarefas, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/tarefas", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var tarefas []Tarefas
	if erro = json.NewDecoder(response.Body).Decode(&tarefas); erro != nil {
		canal <- nil
		return
	}

	canal <- tarefas
}

func BuscaEquipesDoUsuario(canal chan<- []Equipes, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/equipes", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var equipes []Equipes
	if erro = json.NewDecoder(response.Body).Decode(&equipes); erro != nil {
		canal <- nil
		return
	}

	canal <- equipes
}
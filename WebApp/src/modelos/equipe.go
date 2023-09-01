package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

type Equipes struct {
	Id            uint64 `json:"id,omitempty"`
	Nome          string `json:"nome,omitempty"`
	Descricao     string `json:"descricao,omitempty"`
	AutorId       uint64 `json:"autorId,omitempty"`
	Participantes []Usuario
	Tarefas       []Tarefas
}

// BuscarEquipeCompleta chama 3 funções para trazer todas as informações da equipe
func BuscarEquipeCompleta(equipeId uint64, r *http.Request) (Equipes, error) {
	canalEquipe := make(chan Equipes)
	canalUsuarios := make(chan []Usuario)
	canalTarefas := make(chan []Tarefas)

	go BuscaEquipe(canalEquipe, equipeId, r)
	go BuscaUsuariosDaEquipe(canalUsuarios, equipeId, r)
	go BuscaTarefasDaEquipe(canalTarefas, equipeId, r)

	var (
		equipeCompleta Equipes
		tarefas        []Tarefas
		usuarios       []Usuario
		equipeErro     error
	)

	for i := 0; i < 3; i++ {
		select {
		case equipeCarregada := <-canalEquipe:
			if equipeCarregada.Id == 0 {
				equipeErro = errors.New("Erro ao buscar Equipe")
			}
			equipeCompleta = equipeCarregada

		case usuariosCarregados := <-canalUsuarios:
			if usuariosCarregados == nil {
				usuarios = []Usuario{}
			} else {
				usuarios = usuariosCarregados
			}

		case tarefasCarregadas := <-canalTarefas:
			if tarefasCarregadas == nil {
				tarefas = []Tarefas{}
			} else {
				tarefas = tarefasCarregadas
			}
		}

	}
	if equipeErro != nil {
		return Equipes{}, nil
	}

	equipeCompleta.Participantes = usuarios
	equipeCompleta.Tarefas = tarefas

	return equipeCompleta, nil

}

// BuscaEquipe busca as informações da Equipe
func BuscaEquipe(canal chan<- Equipes, equipeId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/equipes/%d", config.APIURL, equipeId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Equipes{}
		return
	}
	defer response.Body.Close()

	var equipe Equipes
	if erro = json.NewDecoder(response.Body).Decode(&equipe); erro != nil {
		canal <- Equipes{}
		return
	}

	canal <- equipe
}

// BuscaUsuariosDaEquipe busca na api os participantes da equipe
func BuscaUsuariosDaEquipe(canal chan<- []Usuario, equipeId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/equipes/%d/usuarios", config.APIURL, equipeId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- []Usuario{}
		return
	}
	defer response.Body.Close()

	var usuarios []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		canal <- []Usuario{}
		return
	}

	canal <- usuarios
}

// BuscaTarefasDaEquipe busca na api as tarefas da equipe
func BuscaTarefasDaEquipe(canal chan<- []Tarefas, equipeId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/equipes/%d/tarefas", config.APIURL, equipeId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- []Tarefas{}
		return
	}
	defer response.Body.Close()

	var tarefas []Tarefas
	if erro = json.NewDecoder(response.Body).Decode(&tarefas); erro != nil {
		canal <- []Tarefas{}
		return
	}

	canal <- tarefas
}

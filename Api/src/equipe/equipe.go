package equipe

import (
	"api/src/modelos"
	"errors"
	"strings"
)

type Equipes struct {
	Id            uint64             `json:"id,omitempty"`
	Nome          string             `json:"nome,omitempty"`
	Descricao     string             `json:"descricao,omitempty"`
	AutorId       uint64             `json:"autorId,omitempty"`
	Participantes []modelos.Usuarios `json:"participantes,omitempty"`
	Tarefas       []modelos.Tarefas  `json:"tarefas,omitempty"`
}

func (equipe *Equipes) Preparar() error {
	if erro := equipe.validar(); erro != nil {
		return erro
	}

	equipe.formatar()

	return nil
}

func (equipe *Equipes) validar() error {
	if equipe.Nome == "" {
		return errors.New("o campo nome não pode estar em branco!!")
	}

	return nil
}

func (equipe *Equipes) formatar() {
	equipe.Nome = strings.TrimSpace(equipe.Nome)
	equipe.Descricao = strings.TrimSpace(equipe.Descricao)
}

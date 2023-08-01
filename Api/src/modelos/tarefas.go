package modelos

import (
	"errors"
	"strings"
)

type Tarefas struct {
	Id        uint64 `json:"id,omitempty"`
	Tarefa    string `json:"tarefa,omitempty"`
	Obsevacao string `json:"observacao,omitempty"`
	AutorId   uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"autorNick,omitempty"`
	Prazo     string `json:"prazo,omitempty"`
}

func (tarefas *Tarefas) Preparar() error {
	if erro := tarefas.validar(); erro != nil {
		return erro
	}

	tarefas.formatar()

	return nil
}

func (tarefas *Tarefas) validar() error {
	if tarefas.Tarefa == "" {
		return errors.New("o campo tarefa não pode estar em branco!!")
	}

	if tarefas.Prazo == "" {
		return errors.New("o campo prazo deve conter uma data")
	}

	return nil
}

func (tarefas *Tarefas) formatar() {
	tarefas.Tarefa = strings.TrimSpace(tarefas.Tarefa)
	tarefas.Obsevacao = strings.TrimSpace(tarefas.Obsevacao)
	tarefas.Prazo = strings.TrimSpace(tarefas.Prazo)
}

package equipe

import "api/src/modelos"

type Equipes struct {
	Id            uint64             `json:"id,omitempty"`
	Nome          string             `json:"nome,omitempty"`
	Descricao     string             `json:"descricao"`
	AutorId       uint64             `json:"autorId,omitempty"`
	Participantes []modelos.Usuarios `json:"participantes"`
	Tarefas       []modelos.Tarefas  `json:"tarefas"`
}

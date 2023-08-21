package modelos

type Tarefas struct {
	Id        uint64 `json:"id,omitempty"`
	Tarefa    string `json:"tarefa,omitempty"`
	Obsevacao string `json:"observacao,omitempty"`
	AutorId   uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"autorNick,omitempty"`
	Prazo     string `json:"prazo,omitempty"`
}

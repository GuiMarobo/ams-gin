package domain

type Turma struct {
	ID         string    `json:"id"`
	Nome       string    `json:"nome"`
	Disciplina string    `json:"disciplina"`
	Professor  string    `json:"professor"`
	AlunoIDs   []string  `json:"-"`
	Alocacao   *Alocacao `json:"alocacao"`
}

type Alocacao struct {
	SalaID        string `json:"sala_id"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}

package domain

type Turma struct {
	ID         string    `json:"id" binding:"required"`
	Nome       string    `json:"nome" binding:"required"`
	Disciplina string    `json:"disciplina" binding:"required"`
	Professor  string    `json:"professor" binding:"required"`
	AlunoIDs   []string  `json:"-"`
	Alocacao   *Alocacao `json:"-"`
}

type Alocacao struct {
	SalaID        string `json:"sala_id" binding:"required"`
	DiaSemana     string `json:"dia_semana" binding:"required,oneof=segunda terca quarta quinta sexta sabado domingo"`
	HorarioInicio string `json:"horario_inicio" binding:"required,len=5,datetime=15:04"`
	HorarioFim    string `json:"horario_fim" binding:"required,len=5,datetime=15:04"`
}

type TurmaResumo struct {
	ID               string    `json:"id"`
	Nome             string    `json:"nome"`
	Disciplina       string    `json:"disciplina"`
	Professor        string    `json:"professor"`
	QuantidadeAlunos int       `json:"quantidade_alunos"`
	Alocada          bool      `json:"alocada"`
	Alocacao         *Alocacao `json:"alocacao"`
}

var OrdemDia = map[string]int{
	"segunda": 1,
	"terca": 2,
	"quarta": 3,
	"quinta": 4,
	"sexta": 5,
	"sabado": 6,
	"domingo": 7,
}

func (a Alocacao) ConflitaCom(b Alocacao) bool {
	return a.DiaSemana == b.DiaSemana &&
		a.HorarioInicio < b.HorarioFim &&
		a.HorarioFim > b.HorarioInicio
}

func (t Turma) Resumo() TurmaResumo {
	return TurmaResumo{
		ID:               t.ID,
		Nome:             t.Nome,
		Disciplina:       t.Disciplina,
		Professor:        t.Professor,
		QuantidadeAlunos: len(t.AlunoIDs),
		Alocada:          t.Alocacao != nil,
		Alocacao:         t.Alocacao,
	}
}

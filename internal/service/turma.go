package service

import (
	"fmt"
	"slices"
	"sort"

	"api-gin/internal/domain"
)

func (s *Servico) CriarTurma(turma domain.Turma) (domain.TurmaResumo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.turmas[turma.ID]; existe {
		return domain.TurmaResumo{}, fmt.Errorf("%w: turma %s já cadastrada", domain.ErrDuplicado, turma.ID)
	}

	s.turmas[turma.ID] = &turma
	return turma.Resumo(), nil
}

func (s *Servico) ListarTurmas() []domain.TurmaResumo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	lista := []domain.TurmaResumo{}
	for _, turma := range s.turmas {
		lista = append(lista, turma.Resumo())
	}
	sort.Slice(lista, func(i, j int) bool { return lista[i].ID < lista[j].ID })
	return lista
}

func (s *Servico) ListarAlunosDaTurma(turmaID string) ([]domain.Aluno, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	turma, existe := s.turmas[turmaID]
	if !existe {
		return nil, fmt.Errorf("%w: turma %s", domain.ErrNaoEncontrado, turmaID)
	}

	lista := []domain.Aluno{}
	for _, alunoID := range turma.AlunoIDs {
		lista = append(lista, *s.alunos[alunoID])
	}
	return lista, nil
}

func (s *Servico) Matricular(turmaID, alunoID string) (domain.TurmaResumo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	turma, existe := s.turmas[turmaID]
	if !existe {
		return domain.TurmaResumo{}, fmt.Errorf("%w: turma %s", domain.ErrNaoEncontrado, turmaID)
	}
	if _, existe := s.alunos[alunoID]; !existe {
		return domain.TurmaResumo{}, fmt.Errorf("%w: aluno %s", domain.ErrNaoEncontrado, alunoID)
	}
	if slices.Contains(turma.AlunoIDs, alunoID) {
		return domain.TurmaResumo{}, fmt.Errorf("%w: aluno %s já matriculado na turma %s", domain.ErrDuplicado, alunoID, turmaID)
	}

	if turma.Alocacao != nil {
		sala := s.salas[turma.Alocacao.SalaID]
		if len(turma.AlunoIDs)+1 > sala.Capacidade {
			return domain.TurmaResumo{}, fmt.Errorf("%w: sala %s comporta %d alunos", domain.ErrCapacidade, sala.ID, sala.Capacidade)
		}
		if err := s.verificarAgendaDoAluno(alunoID, turmaID, *turma.Alocacao); err != nil {
			return domain.TurmaResumo{}, err
		}
	}

	turma.AlunoIDs = append(turma.AlunoIDs, alunoID)
	return turma.Resumo(), nil
}

func (s *Servico) Alocar(turmaID string, nova domain.Alocacao) (domain.TurmaResumo, error) {
	if nova.HorarioInicio >= nova.HorarioFim {
		return domain.TurmaResumo{}, fmt.Errorf("%w: o início deve ser antes do fim", domain.ErrHorarioInvalido)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	turma, existe := s.turmas[turmaID]
	if !existe {
		return domain.TurmaResumo{}, fmt.Errorf("%w: turma %s", domain.ErrNaoEncontrado, turmaID)
	}
	sala, existe := s.salas[nova.SalaID]
	if !existe {
		return domain.TurmaResumo{}, fmt.Errorf("%w: sala %s", domain.ErrNaoEncontrado, nova.SalaID)
	}

	if sala.Capacidade < len(turma.AlunoIDs) {
		return domain.TurmaResumo{}, fmt.Errorf("%w: sala %s comporta %d alunos e a turma tem %d",
			domain.ErrCapacidade, sala.ID, sala.Capacidade, len(turma.AlunoIDs))
	}

	for _, outra := range s.turmas {
		if outra.ID == turmaID || outra.Alocacao == nil || outra.Alocacao.SalaID != nova.SalaID {
			continue
		}
		if nova.ConflitaCom(*outra.Alocacao) {
			return domain.TurmaResumo{}, fmt.Errorf("%w: sala %s já está ocupada pela turma %s nesse horário",
				domain.ErrConflitoAgenda, sala.ID, outra.ID)
		}
	}

	for _, alunoID := range turma.AlunoIDs {
		if err := s.verificarAgendaDoAluno(alunoID, turmaID, nova); err != nil {
			return domain.TurmaResumo{}, err
		}
	}

	turma.Alocacao = &nova
	return turma.Resumo(), nil
}

func (s *Servico) verificarAgendaDoAluno(alunoID, turmaID string, horario domain.Alocacao) error {
	for _, outra := range s.turmas {
		if outra.ID == turmaID || outra.Alocacao == nil || !slices.Contains(outra.AlunoIDs, alunoID) {
			continue
		}
		if horario.ConflitaCom(*outra.Alocacao) {
			return fmt.Errorf("%w: aluno %s já tem aula na turma %s nesse horário",
				domain.ErrConflitoAgenda, alunoID, outra.ID)
		}
	}
	return nil
}

package service

import (
	"fmt"
	"sort"

	"api-gin/internal/domain"
)

func (s *Servico) CriarSala(sala domain.Sala) (domain.Sala, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.salas[sala.ID]; existe {
		return domain.Sala{}, fmt.Errorf("%w: sala %s já cadastrada", domain.ErrDuplicado, sala.ID)
	}
	if sala.Recursos == nil {
		sala.Recursos = []string{}
	}

	s.salas[sala.ID] = &sala
	return sala, nil
}

func (s *Servico) ListarSalas() []domain.Sala {
	s.mu.RLock()
	defer s.mu.RUnlock()

	lista := []domain.Sala{}
	for _, sala := range s.salas {
		lista = append(lista, *sala)
	}
	sort.Slice(lista, func(i, j int) bool { return lista[i].ID < lista[j].ID })
	return lista
}

func (s *Servico) AgendaDaSala(salaID, dia string) ([]domain.TurmaResumo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, existe := s.salas[salaID]; !existe {
		return nil, fmt.Errorf("%w: sala %s", domain.ErrNaoEncontrado, salaID)
	}

	agenda := []domain.TurmaResumo{}
	for _, turma := range s.turmas {
		if turma.Alocacao == nil || turma.Alocacao.SalaID != salaID {
			continue
		}
		if dia != "" && turma.Alocacao.DiaSemana != dia {
			continue
		}
		agenda = append(agenda, turma.Resumo())
	}

	sort.Slice(agenda, func(i, j int) bool {
		a, b := agenda[i].Alocacao, agenda[j].Alocacao
		if a.DiaSemana != b.DiaSemana {
			return domain.OrdemDia[a.DiaSemana] < domain.OrdemDia[b.DiaSemana]
		}
		return a.HorarioInicio < b.HorarioInicio
	})
	return agenda, nil
}

package service

import (
	"fmt"
	"sort"

	"api-gin/internal/domain"
)

func (s *Servico) CriarAluno(aluno domain.Aluno) (domain.Aluno, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, existe := s.alunos[aluno.ID]; existe {
		return domain.Aluno{}, fmt.Errorf("%w: aluno %s já cadastrado", domain.ErrDuplicado, aluno.ID)
	}

	s.alunos[aluno.ID] = &aluno
	return aluno, nil
}

func (s *Servico) ListarAlunos() []domain.Aluno {
	s.mu.RLock()
	defer s.mu.RUnlock()

	lista := []domain.Aluno{}
	for _, aluno := range s.alunos {
		lista = append(lista, *aluno)
	}
	sort.Slice(lista, func(i, j int) bool { return lista[i].ID < lista[j].ID })
	return lista
}

func (s *Servico) BuscarAluno(id string) (domain.Aluno, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	aluno, existe := s.alunos[id]
	if !existe {
		return domain.Aluno{}, fmt.Errorf("%w: aluno %s", domain.ErrNaoEncontrado, id)
	}
	return *aluno, nil
}

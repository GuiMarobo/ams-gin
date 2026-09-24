package service

import (
	"sync"

	"api-gin/internal/domain"
)

type Servico struct {
	mu     sync.RWMutex
	salas  map[string]*domain.Sala
	alunos map[string]*domain.Aluno
	turmas map[string]*domain.Turma
}

func Novo() *Servico {
	return &Servico{
		salas:  make(map[string]*domain.Sala),
		alunos: make(map[string]*domain.Aluno),
		turmas: make(map[string]*domain.Turma),
	}
}

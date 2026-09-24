package store

import (
	"sync"

	"api-gin/internal/domain"
)

type Dados struct {
	Salas  map[string]*domain.Sala
	Alunos map[string]*domain.Aluno
	Turmas map[string]*domain.Turma
}

type Memoria struct {
	mu    sync.RWMutex
	dados Dados
}

func NovaMemoria() *Memoria {
	return &Memoria{
		dados: Dados{
			Salas:  make(map[string]*domain.Sala),
			Alunos: make(map[string]*domain.Aluno),
			Turmas: make(map[string]*domain.Turma),
		},
	}
}

func (m *Memoria) Ler(fn func(d *Dados) error) error {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return fn(&m.dados)
}

func (m *Memoria) Escrever(fn func(d *Dados) error) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	return fn(&m.dados)
}

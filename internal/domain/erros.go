package domain

import "errors"

var (
	ErrNaoEncontrado   = errors.New("não encontrado")
	ErrDuplicado       = errors.New("duplicidade")
	ErrCapacidade      = errors.New("capacidade insuficiente")
	ErrConflitoAgenda  = errors.New("conflito de agenda")
	ErrHorarioInvalido = errors.New("horário inválido")
)

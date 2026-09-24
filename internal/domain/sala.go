package domain

type Sala struct {
	ID         string   `json:"id" binding:"required"`
	Nome       string   `json:"nome" binding:"required"`
	Capacidade int      `json:"capacidade" binding:"required,gt=0"`
	Recursos   []string `json:"recursos"`
}

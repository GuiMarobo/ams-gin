package handler

import (
	"net/http"

	"api-gin/internal/domain"
	"api-gin/internal/service"

	"github.com/gin-gonic/gin"
)

type TurmaHandler struct {
	svc *service.Servico
}

func NovoTurmaHandler(svc *service.Servico) *TurmaHandler {
	return &TurmaHandler{svc: svc}
}

func (h *TurmaHandler) CriarTurma(c *gin.Context) {
	var turma domain.Turma
	if err := c.ShouldBindJSON(&turma); err != nil {
		responderDadosInvalidos(c, err)
		return
	}

	criada, err := h.svc.CriarTurma(turma)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, criada)
}

func (h *TurmaHandler) ListarTurmas(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.ListarTurmas())
}

func (h *TurmaHandler) MatricularAluno(c *gin.Context) {
	var req struct {
		AlunoID string `json:"aluno_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		responderDadosInvalidos(c, err)
		return
	}

	turma, err := h.svc.Matricular(c.Param("id"), req.AlunoID)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, turma)
}

func (h *TurmaHandler) ListarAlunos(c *gin.Context) {
	alunos, err := h.svc.ListarAlunosDaTurma(c.Param("id"))
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, alunos)
}

func (h *TurmaHandler) AlocarSala(c *gin.Context) {
	var alocacao domain.Alocacao
	if err := c.ShouldBindJSON(&alocacao); err != nil {
		responderDadosInvalidos(c, err)
		return
	}

	turma, err := h.svc.Alocar(c.Param("id"), alocacao)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, turma)
}

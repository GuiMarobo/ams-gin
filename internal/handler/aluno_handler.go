package handler

import (
	"net/http"

	"api-gin/internal/domain"
	"api-gin/internal/service"

	"github.com/gin-gonic/gin"
)

type AlunoHandler struct {
	svc *service.Servico
}

func NovoAlunoHandler(svc *service.Servico) *AlunoHandler {
	return &AlunoHandler{svc: svc}
}

func (h *AlunoHandler) CriarAluno(c *gin.Context) {
	var aluno domain.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		responderDadosInvalidos(c, err)
		return
	}

	criado, err := h.svc.CriarAluno(aluno)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, criado)
}

func (h *AlunoHandler) ListarAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.ListarAlunos())
}

func (h *AlunoHandler) BuscarAluno(c *gin.Context) {
	aluno, err := h.svc.BuscarAluno(c.Param("id"))
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, aluno)
}

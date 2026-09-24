package handler

import (
	"net/http"

	"api-gin/internal/domain"
	"api-gin/internal/service"

	"github.com/gin-gonic/gin"
)

type SalaHandler struct {
	svc *service.Servico
}

func NovoSalaHandler(svc *service.Servico) *SalaHandler {
	return &SalaHandler{svc: svc}
}

func (h *SalaHandler) CriarSala(c *gin.Context) {
	var sala domain.Sala
	if err := c.ShouldBindJSON(&sala); err != nil {
		responderDadosInvalidos(c, err)
		return
	}

	criada, err := h.svc.CriarSala(sala)
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusCreated, criada)
}

func (h *SalaHandler) ListarSalas(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.ListarSalas())
}

func (h *SalaHandler) Agenda(c *gin.Context) {
	salaID := c.Param("id")

	agenda, err := h.svc.AgendaDaSala(salaID, c.Query("dia"))
	if err != nil {
		responderErro(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"sala_id": salaID, "turmas": agenda})
}

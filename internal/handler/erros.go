package handler

import (
	"errors"
	"net/http"

	"api-gin/internal/domain"

	"github.com/gin-gonic/gin"
)

func responderErro(c *gin.Context, err error) {
	status := http.StatusInternalServerError

	switch {
	case errors.Is(err, domain.ErrHorarioInvalido):
		status = http.StatusBadRequest
	case errors.Is(err, domain.ErrNaoEncontrado):
		status = http.StatusNotFound
	case errors.Is(err, domain.ErrDuplicado), errors.Is(err, domain.ErrConflitoAgenda):
		status = http.StatusConflict
	case errors.Is(err, domain.ErrCapacidade):
		status = http.StatusUnprocessableEntity
	}

	c.JSON(status, gin.H{"erro": err.Error()})
}

func responderDadosInvalidos(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos: " + err.Error()})
}

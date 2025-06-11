package inscripcion

import (
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /inscripciones/:usuarioID
func GetUserInscripcion(ctx *gin.Context) {
	userIDStr := ctx.Param("usuarioID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	inscripciones, apiErr := services.InscripcionService.GetInscripcionesByUserId(userID)
	if apiErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": apiErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inscripciones)
}

// POST /inscripciones
func CreateInscripcion(ctx *gin.Context) {
	var inscripcionDto dto.InscripcionDto
	if err := ctx.ShouldBindJSON(&inscripcionDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos o faltantes"})
		return
	}

	created, apiErr := services.InscripcionService.CreateInscripcion(inscripcionDto)
	if apiErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": apiErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

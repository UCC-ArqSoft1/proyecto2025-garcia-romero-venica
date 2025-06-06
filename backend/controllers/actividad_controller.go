package actividad

import (
	"net/http"
	"strconv"

	"proyecto2025-garcia-romero-venica/api/backend/domain"
	"proyecto2025-garcia-romero-venica/api/backend/services/actividad_service"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	actividades, err := actividad_service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las actividades"})
		return
	}
	c.JSON(http.StatusOK, actividades)
}

func GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	actividad, err := actividad_service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	c.JSON(http.StatusOK, actividad)
}

func Create(c *gin.Context) {
	var nuevaActividad domain.Actividad
	if err := c.ShouldBindJSON(&nuevaActividad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := actividad_service.Create(nuevaActividad); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la actividad"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Actividad creada correctamente"})
}

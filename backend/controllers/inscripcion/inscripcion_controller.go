package controller

import (
	"net/http"
	"strconv"
	"time"

	"backend/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InscribirUsuario(c *gin.Context, db *gorm.DB) {
	var input struct {
		UsuarioID   int `json:"usuario_id"`
		ActividadID int `json:"actividad_id"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	var actividad domain.Actividad
	if err := db.First(&actividad, input.ActividadID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	if actividad.Cupo <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No hay cupo disponible"})
		return
	}

	inscripcion := domain.Inscripcion{
		UsuarioID:   input.UsuarioID,
		ActividadID: input.ActividadID,
		Fecha:       time.Now(),
	}
	if err := db.Create(&inscripcion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo inscribir"})
		return
	}

	actividad.Cupo -= 1
	db.Save(&actividad)

	c.JSON(http.StatusCreated, inscripcion)
}

func GetInscripcionesPorUsuario(c *gin.Context, db *gorm.DB) {
	usuarioID, _ := strconv.Atoi(c.Param("usuario_id"))
	var inscripciones []domain.Inscripcion
	if err := db.Preload("Actividad").Where("usuario_id = ?", usuarioID).Find(&inscripciones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener inscripciones"})
		return
	}
	c.JSON(http.StatusOK, inscripciones)
}

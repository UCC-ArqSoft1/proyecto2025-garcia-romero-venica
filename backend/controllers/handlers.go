package client

import (
	"net/http"
	"strconv"
	"time"

	"backend/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetActividadesDisponibles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var actividades []domain.Actividad
		if err := db.Preload("Categoria").Where("estado = ?", "disponible").Find(&actividades).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, actividades)
	}
}

func GetActividadByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var actividad domain.Actividad
		if err := db.Preload("Categoria").First(&actividad, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
			return
		}
		c.JSON(http.StatusOK, actividad)
	}
}

func GetActividadesInscripto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioID, _ := strconv.Atoi(c.Param("id"))
		var inscripciones []domain.Inscripcion
		if err := db.Preload("Actividad.Categoria").Where("usuario_id = ?", usuarioID).Find(&inscripciones).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var actividades []domain.Actividad
		for _, i := range inscripciones {
			var act domain.Actividad
			db.First(&act, i.ActividadID)
			actividades = append(actividades, act)
		}

		c.JSON(http.StatusOK, actividades)
	}
}

func InscribirseActividad(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			UsuarioID   uint `json:"usuario_id"`
			ActividadID uint `json:"actividad_id"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var actividad domain.Actividad
		if err := db.First(&actividad, input.ActividadID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
			return
		}
		if actividad.Cupo <= 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "No hay cupo disponible"})
			return
		}

		inscripcion := domain.Inscripcion{
			UsuarioID:   input.UsuarioID,
			ActividadID: input.ActividadID,
			Fecha:       time.Now(),
		}
		if err := db.Create(&inscripcion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		actividad.Cupo -= 1
		db.Save(&actividad)

		c.JSON(http.StatusCreated, inscripcion)
	}
}

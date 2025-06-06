package controller

import (
	"net/http"
	"strconv"

	"backend/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllActividades(c *gin.Context, db *gorm.DB) {
	var actividades domain.Actividades
	if err := db.Find(&actividades).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener actividades"})
		return
	}
	c.JSON(http.StatusOK, actividades)
}

func GetActividadByID(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	var actividad domain.Actividad
	if err := db.First(&actividad, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}
	c.JSON(http.StatusOK, actividad)
}

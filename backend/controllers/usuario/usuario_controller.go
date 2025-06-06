package controller

import (
	"net/http"
	"strconv"

	"backend/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsuarios(c *gin.Context, db *gorm.DB) {
	var usuarios domain.Usuarios
	if err := db.Find(&usuarios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func GetUsuarioByID(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	var usuario domain.Usuario
	if err := db.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

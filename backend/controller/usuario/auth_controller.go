package usuario

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	usuario, err := services.Login(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	rol := 0
	if usuario.TipoUsuario == "admin" {
		rol = 1
	} else if usuario.TipoUsuario == "socio" {
		rol = 2
	}

	token, err := utils.GenerarJWT(usuario.ID, rol)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userID": usuario.ID,
		"rol":    usuario.TipoUsuario,
	})
}
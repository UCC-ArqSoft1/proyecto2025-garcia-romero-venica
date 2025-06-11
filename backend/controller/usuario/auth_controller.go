package usuario

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	// Estructura para recibir los datos del login
	var loginData struct {
		Email    string `json:"username"` // desde el frontend llega como 'username'
		Password string `json:"password"`
	}

	// Validar JSON recibido
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar credenciales (email y contraseña)
	usuario, err := services.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Generar token JWT con ID y rol
	token, err := utils.GenerarJWT(usuario.ID, usuario.TipoUsuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
		return
	}

	// Enviar respuesta con token y datos básicos del usuario
	ctx.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userID": usuario.ID,
		"rol":    usuario.TipoUsuario,
	})
}

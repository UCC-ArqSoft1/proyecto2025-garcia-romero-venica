package app

import (
	"github.com/gin-gonic/gin"

	// Controladores
	"backend/controllers/actividad"
	"backend/controllers/usuario"
	"backend/controllers/inscripcion"

	// Middleware
	"backend/middleware"
)

func MapUrls(router *gin.Engine) {
	api := router.Group("/api")

	// Rutas para actividades deportivas
	api.GET("/actividades", actividad.GetAll)
	api.GET("/actividades/:id", actividad.GetByID)
	api.POST("/actividades", middleware.AuthMiddleware("admin"), actividad.Create)

	// Rutas para inscripciones
	api.GET("/inscripciones", middleware.AuthMiddleware("admin"), inscripcion.GetAll)
	api.POST("/inscripciones", middleware.AuthMiddleware("socio"), inscripcion.Create)
	api.GET("/inscripciones/usuario", middleware.AuthMiddleware("socio"), inscripcion.GetByUsuarioID)

	// Rutas para usuarios
	api.POST("/login", usuario.Login)
	api.GET("/usuarios", usuario.GetAll)
	api.POST("/usuarios", usuario.Create)
}

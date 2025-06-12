package app

import (
	actividadController "backend/controller/actividad"
	inscripcionController "backend/controller/inscripcion"
	authController "backend/controller/usuario"
	"backend/middleware"

	log "github.com/sirupsen/logrus"
)

func mapsUrls() {
	log.Info("Starting mappings configurations")

	// Rutas públicas
	router.POST("/login", authController.Login)
	router.GET("/actividades", actividadController.GetAll)
	router.GET("/actividades/:id", actividadController.GetActividadByID)

	// Rutas protegidas con middleware JWT
	api := router.Group("/")
	api.Use(middleware.AuthMiddleware())

	// Actividades protegidas (ej: creación solo para admin)
	api.POST("/actividades", actividadController.Create)
	api.PUT("/actividades/:id", actividadController.Update)
	api.DELETE("/actividades/:id", actividadController.Delete)
	api.DELETE("/inscripciones/:id", inscripcionController.DeleteInscripcion)


	// Inscripciones protegidas
	api.POST("/inscripciones", inscripcionController.CreateInscripcion)
	api.GET("/inscripciones/:usuarioID", inscripcionController.GetUserInscripcion)
}

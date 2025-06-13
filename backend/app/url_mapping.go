package app


import (
	actividadController "backend/controller/actividad"
	inscripcionController "backend/controller/inscripcion"
	log "github.com/sirupsen/logrus"
	authController "backend/controller/usuario"
	"backend/middleware"
	

	


)
func mapUrls() {


	log.Info("Starting mappings configurations")
	router.POST("/login", authController.Login)
	router.GET("/actividades", actividadController.GetAll)
	router.GET("/actividades/:id", actividadController.GetActividadByID)


	// Actividad


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


package app


import (
	actividadController "backend/controller/actividad"
	inscripcionController "backend/controller/inscripcion"
	log "github.com/sirupsen/logrus"
)


func mapsUrls(){


	log.Info("Starting mappings configurations")
	
	//son acciones que tienen que ver solamente con actividades
	router.GET("/actividades",actividadController.GetAll)
	router.GET("/actividades/:id",actividadController.GetActividadByID)
	router.POST("/actividades", actividadController.Create)
 	router.PUT("/actividades/:id", actividadController.Update)
	router.DELETE("/actividades/:id", actividadController.Delete)
	
	//acciones que tienen que ver con inscripciones 
	router.POST("/inscripciones",inscripcionController.CreateInscripcion)
	router.GET("/inscripciones/:usuarioID",inscripcionController.GetUserInscripcion)


}
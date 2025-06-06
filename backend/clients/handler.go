package client

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	actividadController "backend/controllers/actividad"
	inscripcionController "backend/controllers/inscripcion"
)

func RegisterSocioRoutes(r *gin.Engine, db *gorm.DB) {
	socio := r.Group("/socio")
	{
		socio.GET("/actividades", func(c *gin.Context) {
			actividadController.GetAllActividades(c, db)
		})
		socio.GET("/actividades/:id", func(c *gin.Context) {
			actividadController.GetActividadByID(c, db)
		})
		socio.GET("/inscripciones/:usuario_id", func(c *gin.Context) {
			inscripcionController.GetInscripcionesPorUsuario(c, db)
		})
		socio.POST("/inscribirse", func(c *gin.Context) {
			inscripcionController.InscribirUsuario(c, db)
		})
	}
}

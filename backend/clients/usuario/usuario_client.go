package usuario

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	inscripcionController "backend/controllers/inscripcion"
)

func RegisterUsuarioRoutes(r *gin.Engine, db *gorm.DB) {
	usuario := r.Group("/usuario")
	{
		usuario.GET("/inscripciones/:usuario_id", func(c *gin.Context) {
			inscripcionController.GetInscripcionesPorUsuario(c, db)
		})
		usuario.POST("/inscribirse", func(c *gin.Context) {
			inscripcionController.InscribirUsuario(c, db)
		})
	}
}

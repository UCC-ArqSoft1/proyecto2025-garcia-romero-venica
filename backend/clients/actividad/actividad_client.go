package actividad

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	actividadController "backend/controllers/actividad"
)

func RegisterActividadRoutes(r *gin.Engine, db *gorm.DB) {
	actividad := r.Group("/actividad")
	{
		actividad.GET("/", func(c *gin.Context) {
			actividadController.GetAllActividades(c, db)
		})
		actividad.GET("/:id", func(c *gin.Context) {
			actividadController.GetActividadByID(c, db)
		})
	}
}

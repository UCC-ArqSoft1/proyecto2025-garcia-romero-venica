package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	actividadClient "backend/clients/actividad"
	usuarioClient "backend/clients/usuario"
	"backend/domain"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("gimnasio.db"), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos")
	}

	// Migraciones
	db.AutoMigrate(&domain.Usuario{}, &domain.Actividad{}, &domain.Inscripcion{})

	// Rutas
	usuarioClient.RegisterUsuarioRoutes(r, db)
	actividadClient.RegisterActividadRoutes(r, db)

	r.Run(":8080")
}

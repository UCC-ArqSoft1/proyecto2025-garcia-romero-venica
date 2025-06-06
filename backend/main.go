package main

import (
	"backend/bd"
	"backend/app"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta con la base de datos y ejecuta migraciones
	bd.StartDbEngine()

	// Inicializa el router de Gin
	router := gin.Default()

	// Mapea las rutas de la API
	app.MapUrls(router)

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")
}

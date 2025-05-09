package main

import (
	"log"
	"net/http"

	"github.com/tu_usuario/gimnasio/handler"
	"github.com/tu_usuario/gimnasio/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/login", handler.Login)
	mux.Handle("/api/actividades", middleware.Auth(handler.ActividadesHandler))
	mux.Handle("/api/admin/crear", middleware.Admin(handler.CrearActividad))

	log.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

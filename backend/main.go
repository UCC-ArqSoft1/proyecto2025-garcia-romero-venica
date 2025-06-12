package main
import (
	"backend/app"
	"backend/bd"
	"log"
)
func main (){
	db.Init()
	db.StartDbEngine()
	app.StartRoute()
	log.Println("aplicacion iniciada correctamente")
	
	
}

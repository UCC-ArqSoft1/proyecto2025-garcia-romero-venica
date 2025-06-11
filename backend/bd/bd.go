package db

import (
	actividadClient "backend/clients/actividad"
	inscripcionClient "backend/clients/inscripcion"
	usuarioClient "backend/clients/usuario"
	"backend/domain"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// Configuración para MySQL (según tu dump)
	dsn := "root:Belgrano11@tcp(localhost:3306)/gimnasio?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Error("Connection to database failed")
		log.Fatal(err)
	} else {
		log.Info("Database connection established")
	}

	// Asignación a clients
	actividadClient.DB = DB
	usuarioClient.DB = DB
	inscripcionClient.DB = DB

}

func StartDbEngine() {
	DB.AutoMigrate(&domain.Actividad{})
	DB.AutoMigrate(&domain.Usuario{})
	DB.AutoMigrate(&domain.Inscripcion{})

	log.Info("Finishing Migration Database Tables")
}

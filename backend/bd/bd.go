package db
import (
	actividadClient "backend/clients/actividad"
	usuarioClient "backend/clients/usuario"
	inscripcionClient "backend/clients/inscripcion"
	"gorm.io/driver/mysql"
	"backend/domain"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"

)

var(
	DB *gorm.DB
	err error

)

func Init() {
	// Configuración para MySQL (según tu dump)
dsn := "andrew:123456@tcp(localhost:3306)/gimnasio?parseTime=true"
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
	actividadClient.DB=DB
	usuarioClient.DB =DB
	inscripcionClient.DB=DB


}



func StartDbEngine(){
	DB.AutoMigrate(&domain.Actividad{})
	DB.AutoMigrate(&domain.Usuario{})
	DB.AutoMigrate(&domain.Inscripcion{})
	

	log.Info("Finishing Migration Database Tables")
}

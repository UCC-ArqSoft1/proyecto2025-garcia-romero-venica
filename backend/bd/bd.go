package db
import (
	actividadClient 	"proyecto2025-garcia-romero-venica/clients/actividad"
	usuarioClient   	"proyecto2025-garcia-romero-venica/clients/usuario"
	inscripcionClient	"proyecto2025-garcia-romero-venica/clients/inscripcion"
	"proyecto2025-garcia-romer-venica/api/backend/domain"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	log "github.com/sirupsen/logrus"

)

var(
	dba *gorm.DB
	err error

)
func init(){
	dba, err=gorm.Open(mysql.Open("file::memory?cache=shared"), &gorm.config{
		Logger: logger.Default.LogMode (logger.Info),

	})

	if err != nill{
		log.Info("Connection failed")
		log.fatal(err)

	}else{
		log.info("Connection Established")
	}

	actividadClient.Db= dba
	usuarioClient.Db=   dba
	inscripcionClient=  dba
	categoriaclient=    dba
	tipo_usuarioClient= dba
}
func StartDbEngine(){
	dba.AutoMigrate(&domain.actividad{})
	dba.AutoMigrate(&domain.inscripcion{})
	dba.AutoMigrate(&domain.usuario{})
	log.Info("Finishing Migration Database Tables")
	prePreInsertData()
}

package db

import (
	"log"

	actividadClient "proyecto2025-garcia-romero-venica/clients/actividad"
	usuarioClient "proyecto2025-garcia-romero-venica/clients/usuario"
	inscripcionClient "proyecto2025-garcia-romero-venica/clients/inscripcion"

	actividadService "proyecto2025-garcia-romero-venica/api/backend/services/actividad_service"
	usuarioController "proyecto2025-garcia-romero-venica/api/backend/controllers/usuario"

	"proyecto2025-garcia-romero-venica/api/backend/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dba *gorm.DB
	err error
)

func init() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/gimnasio?parseTime=true"
	dba, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	} else {
		log.Println("Conexión a la base de datos exitosa")
	}

	// Inyección de dependencias
	actividadClient.Db = dba
	usuarioClient.Db = dba
	inscripcionClient.Db = dba

	actividadService.SetDB(dba)
	usuarioController.SetDB(dba)
}

func StartDbEngine() {
	err := dba.AutoMigrate(&domain.Actividad{}, &domain.Inscripcion{}, &domain.Usuario{})
	if err != nil {
		log.Fatal("Error al migrar tablas:", err)
	}
	log.Println("Migración de tablas completada")
}

package inscripcion

import (
	"backend/domain"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetInscripcionesCountByActividad(actividadID int) (int64, error) {
	var count int64

	err := DB.Model(&domain.Inscripcion{}).
		Where("actividades_id = ?", actividadID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserInscriptions(userID int) ([]domain.Inscripcion, error) {
	var inscripciones []domain.Inscripcion

	err := DB.Preload("Actividad").
		Preload("Usuario").
		Where("usuarios_id = ?", userID).
		Find(&inscripciones).Error

	if err != nil {
		log.Error("Error obteniendo inscripciones: ", err)
		return nil, err
	}

	return inscripciones, nil
}

func GetInscripcionByUserAndActivity(userID int, actividadID int) (domain.Inscripcion, error) {
	var inscripcion domain.Inscripcion
	err := DB.Where("usuarios_id = ? AND actividades_id = ?", userID, actividadID).First(&inscripcion).Error
	if err != nil {
		return domain.Inscripcion{}, err
	}
	return inscripcion, nil
}

func CreateInscripcion(inscripcion domain.Inscripcion) (domain.Inscripcion, error) {
	var existing domain.Inscripcion
	err := DB.Where("usuarios_id = ? AND actividades_id = ?",
		inscripcion.UsuarioID, inscripcion.ActividadID).
		First(&existing).Error

	if err == nil {
		return domain.Inscripcion{}, fmt.Errorf("el usuario ya está inscrito en esta actividad")
	} else if err != gorm.ErrRecordNotFound {
		return domain.Inscripcion{}, err
	}

	result := DB.Create(&inscripcion)
	if result.Error != nil {
		log.Error("Error creando inscripcion: ", result.Error)
		return domain.Inscripcion{}, result.Error
	}

	log.Debug("Inscripcion creada: ", inscripcion.ID)
	return inscripcion, nil
}

// función para eliminar inscripción por ID
func DeleteInscripcion(id int) error {
	result := DB.Delete(&domain.Inscripcion{}, id)
	if result.Error != nil {
		log.Error("Error al eliminar inscripción: ", result.Error)
		return result.Error
	}
	log.Info("Inscripción eliminada con ID:", id)
	return nil
}

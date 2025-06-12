package actividad

import (
	"backend/domain"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ✅ Obtener todas las actividades activas
func GetActividades() []domain.Actividad {
	var actividades []domain.Actividad
	DB.Where("estado = ?", true).Find(&actividades)
	return actividades
}

// ✅ Obtener una actividad por ID (corregido: usa id_actividades)
func GetActividadById(id int) (domain.Actividad, error) {
	var actividad domain.Actividad

	// 🔧 correction ici
	result := DB.Where("id_actividades = ? AND estado = true", id).First(&actividad)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Info("Actividad no encontrada con ID:", id)
			return domain.Actividad{}, nil
		}
		log.Error("Error al buscar actividad:", result.Error)
		return domain.Actividad{}, result.Error
	}

	log.Debug("Actividad encontrada:", actividad)
	return actividad, nil
}

// ✅ Insertar nueva actividad
func InsertActividad(actividad domain.Actividad) domain.Actividad {
	result := DB.Create(&actividad)
	if result.Error != nil {
		log.Error(result.Error)
	}
	log.Debug("Actividad creada con ID:", actividad.ID)
	return actividad
}

// ✅ Borrado lógico (cambia estado à false)
func DeleteActividad(id int) (domain.Actividad, error) {
	var actividad domain.Actividad

	// 🔧 utiliser id_actividades ici aussi
	if err := DB.Where("id_actividades = ?", id).First(&actividad).Error; err != nil {
		return domain.Actividad{}, err
	}

	if err := DB.Model(&actividad).Update("estado", false).Error; err != nil {
		return domain.Actividad{}, err
	}

	log.Info("Actividad desactivada, ID:", id)
	return actividad, nil
}

// ✅ Actualizar actividad (correction WHERE id)
func ActualizarActividad(id int, actividad domain.Actividad) (domain.Actividad, error) {
	result := DB.Model(&domain.Actividad{}).
		Where("id_actividades = ?", id).
		Updates(actividad)

	if result.Error != nil {
		log.Error("Error al actualizar actividad:", result.Error)
		return domain.Actividad{}, result.Error
	}

	var updatedActividad domain.Actividad
	if err := DB.Where("id_actividades = ?", id).First(&updatedActividad).Error; err != nil {
		return domain.Actividad{}, err
	}

	log.Debug("Actividad actualizada:", updatedActividad)
	return updatedActividad, nil
}



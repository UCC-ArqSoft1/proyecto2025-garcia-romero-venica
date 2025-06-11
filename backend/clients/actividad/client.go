package actividad
import( 
	"backend/domain"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB




//aqui deberiamos poder obtener todas las actividades con estado 1 
func GetActividades() []domain.Actividad {
	var actividades []domain.Actividad
	DB.Where("estado = ?", true).Find(&actividades)
	return actividades
}

//aqui obtenemos las actividades por medio de un ID y se muestran las actividades con estado 1 
func GetActividadById(id int) (domain.Actividad, error) {
	var actividad domain.Actividad

	result := DB.Where("id = ? AND estado = true", id).First(&actividad)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Info("Actividad no encontrada con ID:", id)
			return domain.Actividad{}, nil // si prefieres que no sea error si no existe
		}
		log.Error("Error al buscar actividad:", result.Error)
		return domain.Actividad{}, result.Error
	}

	log.Debug("Actividad encontrada:", actividad)
	return actividad, nil
}

//se inserta actividad
func InsertActividad(actividad domain.Actividad) domain.Actividad {
    result := DB.Create(&actividad)
    if result.Error != nil {
        log.Error(result.Error)
    }
    log.Debug("Actividad creada con ID:", actividad.ID)
    return actividad
}
//se borra una actividad por medio de un ID, es un borrado logico
func DeleteActividad(id int) (domain.Actividad, error) {
    var actividad domain.Actividad
    if err := DB.First(&actividad, id).Error; err != nil {
        return domain.Actividad{}, err
    }

    if err := DB.Model(&actividad).Update("estado", 0).Error; err != nil {
        return domain.Actividad{}, err
    }

    log.Info("Actividad desactivada, ID:", id)
    return actividad, nil
}



//se obtienen las actividades y podriamos utilizar esto despues de editar una actividad
func ActualizarActividad(id int, actividad domain.Actividad) (domain.Actividad, error) {
    result := DB.Model(&domain.Actividad{}).Where("id= ?", id).Updates(actividad)
    if result.Error != nil {
        log.Error("Error al actualizar actividad:", result.Error)
        return domain.Actividad{}, result.Error
    }
    
    // Obtener la actividad actualizada para devolverla
    var updatedActividad domain.Actividad
    if err := DB.First(&updatedActividad, id).Error; err != nil {
        return domain.Actividad{}, err
    }
    
    log.Debug("Actividad actualizada:", updatedActividad)
    return updatedActividad, nil
}

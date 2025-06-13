package inscripcion
import( 
	"backend/domain"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var DB *gorm.DB


func GetInscripcionesCountByActividad(actividadID int) (int64, error) {
    var count int64

    err := DB.Model(&domain.Inscripcion{}).
        Where("actividad_id = ?", actividadID).
        Count(&count).Error
    if err != nil {
        return 0, err
    }
    return count, nil
}
func DeleteInscripcion(id int) error {
	result := DB.Delete(&domain.Inscripcion{}, id)
	if result.Error != nil {
		log.Error("Error al eliminar inscripción: ", result.Error)
		return result.Error
	}
	log.Info("Inscripción eliminada con ID:", id)
	return nil
}


func GetUserInscriptions(userID int) ([]domain.Inscripcion, error) {
    var inscripciones []domain.Inscripcion
    
    err := DB.Preload("Actividad").
             Preload("Usuario"). // Si necesitas datos del usuario
             Where("usuario_id = ?", userID).
             Find(&inscripciones).
             Error

    if err != nil {
        log.Error("Error obteniendo inscripciones: ", err)
        return nil, err
    }
    
    return inscripciones, nil
}
func GetInscripcionByUserAndActivity(userID int, actividadID int) (domain.Inscripcion, error) {
    var inscripcion domain.Inscripcion
    err := DB.Where("usuario_id = ? AND actividad_id = ?", userID, actividadID).First(&inscripcion).Error
    if err != nil {
        return domain.Inscripcion{}, err
    }
    return inscripcion, nil
}

func CreateInscripcion(inscripcion domain.Inscripcion) (domain.Inscripcion, error) {
    // Verificar si ya existe una inscripción para este usuario en esta actividad
    var existing domain.Inscripcion
    err := DB.Where("usuario_id = ? AND actividad_id = ?", 
                   inscripcion.UsuarioID, inscripcion.ActividadID).
             First(&existing).Error
    
    if err == nil {
	return domain.Inscripcion{}, fmt.Errorf("el usuario ya está inscrito en esta actividad")
    } else if err != gorm.ErrRecordNotFound {
        return domain.Inscripcion{}, err
    }
      var actividad domain.Actividad
    if err := DB.First(&actividad, inscripcion.ActividadID).Error; err != nil {
        return domain.Inscripcion{}, fmt.Errorf("actividad no encontrada")
    }
    
    if actividad.Disponible <= 0 {
        return domain.Inscripcion{}, fmt.Errorf("no hay cupo disponible para esta actividad")
    }
    
    // Disminuir el disponible (actualización directa en la base de datos)
    if err := DB.Model(&domain.Actividad{}).
        Where("id = ?", inscripcion.ActividadID).
        Update("disponible", gorm.Expr("disponible - 1")).Error; err != nil {
        return domain.Inscripcion{}, fmt.Errorf("error actualizando cupo disponible")
    }
    
    result := DB.Create(&inscripcion)
    if result.Error != nil {
        log.Error("Error creando inscripcion: ", result.Error)
        return domain.Inscripcion{}, result.Error
    }
    
    log.Debug("Inscripcion creada: ", inscripcion.ID)
    return inscripcion, nil
}
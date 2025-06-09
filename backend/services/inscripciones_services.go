package services

import (
	inscripcionClient "backend/clients/inscripcion"
	actividadClient "backend/clients/actividad"
	usuarioClient "backend/clients/usuario"
	"backend/dto"
	"backend/domain"
	"time"
	"errors"
)

type inscripcionService struct{}

type inscripcionServiceInterface interface {
	CreateInscripcion(inscripcionDto dto.InscripcionDto) (dto.InscripcionDto, error)
	GetInscripcionesByUserId(userId int) (dto.InscripcionesDto, error)
}

var (
	InscripcionService inscripcionServiceInterface
)

func init() {
	InscripcionService = &inscripcionService{}
}

func (s *inscripcionService) GetInscripcionesByUserId(userId int) (dto.InscripcionesDto, error) {
	inscripciones, err := inscripcionClient.GetUserInscriptions(userId)
	if err != nil {
		return nil, errors.New("error obteniendo inscripciones: " + err.Error())
	}
	var inscripcionesDto dto.InscripcionesDto

	for _, inscripcion := range inscripciones {
		inscripcionDto := dto.InscripcionDto{
			ID:          inscripcion.ID,
			ActividadID: inscripcion.ActividadID,
			UsuarioID:   inscripcion.UsuarioID,
			Fecha:       inscripcion.Fecha,
		}

		inscripcionesDto = append(inscripcionesDto, inscripcionDto)
	}

	return inscripcionesDto, nil
}

func (s *inscripcionService) CreateInscripcion(inscripcionDto dto.InscripcionDto) (dto.InscripcionDto, error) {
	// Verificar si la actividad existe
	actividad, err := actividadClient.GetActividadById(inscripcionDto.ActividadID)
	if err != nil {
		return dto.InscripcionDto{}, errors.New("error buscando actividad: " + err.Error())
	}
	if actividad.ID == 0 {
		return dto.InscripcionDto{}, errors.New("actividad no encontrada")
	}

	// Verificar si el usuario existe
	usuario, err := usuarioClient.GetUserById(inscripcionDto.UsuarioID)
	if err != nil {
		return dto.InscripcionDto{}, errors.New("error buscando usuario: " + err.Error())
	}
	if usuario.ID == 0 {
		return dto.InscripcionDto{}, errors.New("usuario no encontrado")
	}

	// Verificar si ya está inscrito
	existing, err := inscripcionClient.GetInscripcionByUserAndActivity(inscripcionDto.UsuarioID, inscripcionDto.ActividadID)
	if err != nil && err.Error() != "record not found" {
		return dto.InscripcionDto{}, errors.New("error consultando inscripción: " + err.Error())
	}
	if existing.ID != 0 {
		return dto.InscripcionDto{}, errors.New("el usuario ya está inscrito en esta actividad")
	}

	// Verificar disponibilidad
	if actividad.Cupo > 0 {
    count, err := inscripcionClient.GetInscripcionesCountByActividad(inscripcionDto.ActividadID)
    if err != nil {
        return dto.InscripcionDto{}, errors.New("error contando inscripciones: " + err.Error())
    }
    if int(count) >= actividad.Cupo {
        return dto.InscripcionDto{}, errors.New("no hay cupo disponible para esta actividad")
    }
}

	// Crear inscripción
	inscripcion := domain.Inscripcion{
		UsuarioID:   inscripcionDto.UsuarioID,
		ActividadID: inscripcionDto.ActividadID,
		Fecha:       time.Now(),
	}

	created, err := inscripcionClient.CreateInscripcion(inscripcion)
	if err != nil {
		return dto.InscripcionDto{}, errors.New("error creando inscripción: " + err.Error())
	}

	// Preparar respuesta
	response := dto.InscripcionDto{
		ID:          created.ID,
		UsuarioID:   created.UsuarioID,
		ActividadID: created.ActividadID,
		Fecha:       created.Fecha,
	}

	return response, nil
}
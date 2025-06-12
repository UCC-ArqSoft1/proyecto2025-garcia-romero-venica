package services

import (
	"backend/domain"
	"backend/dto"
	actividadClient "backend/clients/actividad"
	log "github.com/sirupsen/logrus"
	"errors"
)

type actividadService struct{}

type actividadServiceInterface interface {
	GetAllActividades() (dto.ActividadesDto, error)
	GetActividadByID(id int) (dto.ActividadDto, error)
	CreateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, error)
	UpdateActividad(id int, actividadDto dto.ActividadDto) (dto.ActividadDto, error)
	DeleteActividad(id int) error
}

var (
	ActividadService actividadServiceInterface
)

func init() {
	ActividadService = &actividadService{}
}

func (s *actividadService) GetAllActividades() (dto.ActividadesDto, error) {
	actividadesDomain := actividadClient.GetActividades()
	if len(actividadesDomain) == 0 {
		return dto.ActividadesDto{}, nil
	}

	var actividadesDto dto.ActividadesDto
	for _, actividad := range actividadesDomain {
		actividadesDto = append(actividadesDto, domainToDto(actividad))
	}

	return actividadesDto, nil
}

func (s *actividadService) GetActividadByID(id int) (dto.ActividadDto, error) {
	actividad, err := actividadClient.GetActividadById(id)
	if err != nil {
		return dto.ActividadDto{}, err // devuelve el error recibido
	}
	if actividad.ID == 0 {
		return dto.ActividadDto{}, errors.New("actividad no encontrada o inactiva")
	}
	return domainToDto(actividad), nil
}

func (s *actividadService) CreateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, error) {
	if actividadDto.Nombre == "" {
		return dto.ActividadDto{}, errors.New("el nombre es requerido")
	}
	if actividadDto.Cupo <= 0 {
		return dto.ActividadDto{}, errors.New("el cupo debe ser mayor a cero")
	}

	newActividad := domain.Actividad{
		Nombre:      actividadDto.Nombre,
		Descripcion: actividadDto.Descripcion,
		Estado:      true,
		Horario:     actividadDto.Horario,
		Cupo:        actividadDto.Cupo,
		Profesor:    actividadDto.Profesor,
		Disponible:  actividadDto.Cupo,
		Categoria:   actividadDto.Categoria,
	}

	createdActividad := actividadClient.InsertActividad(newActividad)
	log.Debug("Actividad creada con ID:", createdActividad.ID)

	return domainToDto(createdActividad), nil
}

func (s *actividadService) UpdateActividad(id int, actividadDto dto.ActividadDto) (dto.ActividadDto, error) {
	existing,err := actividadClient.GetActividadById(id)
	if existing.ID == 0 {
		return dto.ActividadDto{}, errors.New("actividad no encontrada")
	}

	updatedActividad := domain.Actividad{
		Nombre:      actividadDto.Nombre,
		Descripcion: actividadDto.Descripcion,
		Horario:     actividadDto.Horario,
		Cupo:        actividadDto.Cupo,
		Profesor:    actividadDto.Profesor,
		Categoria:   actividadDto.Categoria,
	}

	updatedActividad.Disponible = existing.Disponible + (actividadDto.Cupo - existing.Cupo)
	if updatedActividad.Disponible < 0 {
		return dto.ActividadDto{}, errors.New("no se puede reducir el cupo por debajo de las reservas existentes")
	}

	result, err := actividadClient.ActualizarActividad(id, updatedActividad)
	if err != nil {
		log.Error("Error al actualizar actividad: ", err)
		return dto.ActividadDto{}, err
	}

	return domainToDto(result), nil
}

func (s *actividadService) DeleteActividad(id int) error {
	_, err := actividadClient.DeleteActividad(id)
	if err != nil {
		log.Error("Error al eliminar actividad: ", err)
		return err
	}
	return nil
}

// Función auxiliar para convertir de Domain a DTO
func domainToDto(actividad domain.Actividad) dto.ActividadDto {
	return dto.ActividadDto{
		ID:          actividad.ID,
		Nombre:      actividad.Nombre,
		Descripcion: actividad.Descripcion,
		Estado:      actividad.Estado,
		Cupo:        actividad.Cupo,
		Disponible:  actividad.Disponible,
		Profesor:    actividad.Profesor,
		Horario:     actividad.Horario,
		Categoria:   actividad.Categoria,
	}
}
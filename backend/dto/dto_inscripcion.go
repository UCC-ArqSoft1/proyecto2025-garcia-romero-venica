package dto

import "time"

type InscripcionDto struct {
	ID          int       `json:"id"`
	UsuarioId   int       `json:"usuario_id"`
	Fecha       time.Time `json:"fecha"`
	ActividadID int       `json:"actividad_id"`
}

type InscripcionesDto []InscripcionDto

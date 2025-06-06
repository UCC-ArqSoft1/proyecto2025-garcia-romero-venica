package domain

import "time"

type Inscripcion struct {
	ID          int       `gorm:"primaryKey"`
	UsuarioID   int       `gorm:"column:usuario_id"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID"`
	Fecha       time.Time `gorm:"type:date"`
	ActividadID int       `gorm:"column:actividad_id"`
	Actividad   Actividad `gorm:"foreignKey:ActividadID"`
}

type Inscripciones []Inscripcion

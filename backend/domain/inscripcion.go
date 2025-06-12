package domain

import "time"

type Inscripcion struct {
	ID          int       `gorm:"column:id_inscripciones;primaryKey" json:"id"`
	UsuarioID   int       `gorm:"column:usuarios_id"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID"`
	Fecha       time.Time `gorm:"column:fecha_inscripcion"`
	ActividadID int       `gorm:"column:actividades_id"`
	Actividad   Actividad `gorm:"foreignKey:ActividadID"`
}


func (Inscripcion) TableName() string {
	return "inscripciones"
}


type Inscripciones []Inscripcion

package domain

type Actividad struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Nombre      string `gorm:"type:varchar(45)" json:"nombre"`
	Descripcion string `gorm:"type:varchar(100)" json:"descripcion"`
	Estado      bool   `json:"estado"`
	Horario     string `gorm:"type:varchar(45)" json:"horario"`
	Cupo        int    `json:"cupo"`
	Profesor    string `gorm:"type:varchar(45)" json:"profesor"`
	Disponible  int    `json:"disponible"`
	Categoria   string `gorm:"type:varchar(100)" json:"categoria"`
}

type Actividades []Actividad

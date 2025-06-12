package domain

type Actividad struct {
	ID          int    `gorm:"column:id_actividades;primaryKey" json:"id"`
	Nombre      string `gorm:"type:varchar(45)" json:"nombre"`
	Descripcion string `gorm:"type:varchar(100)" json:"descripcion"`
	Estado      bool   `json:"estado"`
	Horario     string `gorm:"column:horarios" json:"horario"`
	Cupo        int    `json:"cupo"`
	Profesor    string `gorm:"type:varchar(45)" json:"profesor"`
	Disponible  int    `gorm:"column:disponibles" json:"disponible"`
	Categoria   string `gorm:"type:varchar(100)" json:"categoria"`
}

func (Actividad) TableName() string {
	return "actividades"
}

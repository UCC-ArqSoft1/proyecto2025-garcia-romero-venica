package domain
import "time"
type Inscripcion struct {
	ID           int        `gorm:"primaryKey"`
	UsuarioID    int        `gorm:"foreignKey:UsuarioID"`
	Usuario      Usuario    
	Fecha        time.Time 	`gorm:"type:date"`
	ActividadID  int        `gorm:"foreignKey:ActividadID"`
	Actividad    Actividad  
}


type Inscripciones[] Inscripcion



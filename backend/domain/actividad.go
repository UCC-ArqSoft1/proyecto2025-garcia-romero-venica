package domain

type Actividad struct {
	ID          int      `gorm:"primaryKey"`
	Nombre      string   `gorm:"type:varchar(45)"`
	Descripcion string   `gorm:"type:varchar(100)"`
	Estado      bool  
	Horario    	string   `gorm:"type:varchar(45)"`
	Cupo        int		 
	Profesor    string   `gorm:"type:varchar(45)"`
	Disponible  int      
	Categoria 	string   `gorm:"type:varchar(100)"`

}
type Actividades [] Actividad
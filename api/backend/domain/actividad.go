package domain

type Actividad struct {
	ID          int      `gorm:"primaryKey"`
	Nombre      string   `gorm:"type:varchar(45);not null"`
	Descripcion string   `gorm:"type:varchar(100);not null"`
	Estado      string   `gorm:"type:varchar(50)"`
	Cupo        int
	CategoriaID int
	Categoria   Categoria `gorm:"foreignKey:CategoriaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}


type Actividades [] Actividad


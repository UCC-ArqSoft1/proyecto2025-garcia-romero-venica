package domain

type Inscripcion struct {
	ID           int        `gorm:"primaryKey"`
	UsuarioID    int        `gorm:"not null"`
	Usuario      Usuario    `gorm:"foreignKey:UsuarioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Fecha        string  `gorm:"not null"`
	ActividadID  int        `gorm:"not null"`
	Actividad    Actividad  `gorm:"foreignKey:ActividadID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}


type Inscrpciones[] Inscripcion



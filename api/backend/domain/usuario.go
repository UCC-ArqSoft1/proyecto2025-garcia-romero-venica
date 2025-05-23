package domain


type Usuario struct {
	ID             int          `gorm:"primaryKey"`
	Nombre         string       `gorm:"type:varchar(45);not null"`
	Email          string       `gorm:"type:varchar(45);unique;not null"`
	Password       string       `gorm:"type:varchar(100);not null"`
	TipoUsuarioID  int          `gorm:"not null"`
	TipoUsuario    TipoUsuario  `gorm:"foreignKey:TipoUsuarioID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Usuarios [] Usuario

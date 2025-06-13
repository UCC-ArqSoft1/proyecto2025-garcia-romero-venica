package domain


type Usuario struct {
	ID             int          `gorm:"primaryKey"`
	Nombre         string       `gorm:"type:varchar(45);not null"`
	Email          string       `gorm:"type:varchar(45);unique;not null"`
	Password       string       `gorm:"type:varchar(100);not null"`
	TipoUsuario    string      	`gorm:"type:varchar(45)"`
}

type Usuarios [] Usuario

package domain

type TipoUsuario struct {
	ID     int    `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(45);not null"`
}
type TipoUSuarios [] TipoUsuario

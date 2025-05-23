package domain

type Categoria struct {
	ID          int    `gorm:"primaryKey"`
	Nombre      string `gorm:"type:varchar(45);not null"`
	Descripcion string `gorm:"type:varchar(100)"`
}

type Categorias [] Categoria

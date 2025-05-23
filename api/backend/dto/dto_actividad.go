package dto


type ActividadDto struct {
	ID int `json:nombre`
	Nombre string  `json:nombre`
	Descripcion string `json:descripcion`
	Estado bool `json:estado`
	Cupo int `json:cupo`
	Disponible int `json:disponible`
	Profesor string `json:profesor`
	Horario string `json:horario`


	CategoriaID int  `json:categoria_id`
	Categoria Categoria
}

type ActividadesDto [] ActividadDto

package dto

type ActividadDto struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Estado     bool   `json:"estado"`
	Cupo       int    `json:"cupo"`
	Disponible int    `json:"disponible"`
	Profesor   string `json:"profesor"`
	Horario    string `json:"horario"`
	Categoria  string `json:"categoria"`
}

type ActividadesDto []ActividadDto

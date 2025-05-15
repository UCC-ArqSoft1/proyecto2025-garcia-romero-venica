package domain

type Actividad struct {
	Id_actividad int    `json:"id_actividad"`
	Fecha        Fecha  `json:"fecha"`
	Cupo         bool   `json:"cupo"`
	Categoria    string `json:"categoria"`
	Estado       bool   `json:"estado"`
}

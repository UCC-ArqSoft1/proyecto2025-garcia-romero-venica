package domain

type Usuario struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Dni          int    `json:"dni"`
	Usuario      string `json:"usuario"`
	Contraseña   string `json:"contraseña"`
	Telefono     int    `json:"telefono"`
	Tipo_Usuario int    `json:"tipo_usuario"`
}

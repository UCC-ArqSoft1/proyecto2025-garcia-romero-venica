// dto_usuario.go
package dto

type UsuarioDto struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TipoUsuario string `json:"tipo_usuario"`
}

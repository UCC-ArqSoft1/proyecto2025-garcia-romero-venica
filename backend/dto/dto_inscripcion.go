package dto
import (
	"time"
)


type InscripcionDto struct {
	ID 			int 		`json:"id"`
	UsuarioID 	int 		`json:"usuario_id"`
	Fecha 		time.Time	`json:"date"`
	ActividadID int 		`json:"actividad_id"`
}
type InscripcionesDto  []InscripcionDto 
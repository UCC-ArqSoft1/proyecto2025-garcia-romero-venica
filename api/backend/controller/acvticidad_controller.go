package controller

import (
	"encoding/json"
	"net/http"

	"github.com/pablovenica/gimnasio/domain"
	"github.com/pablovenica/gimnasio/utils"
)

var usuarios = []domain.Usuario{
	{ID: 1, Usuario: "admin", Contraseña: "1234", Tipo: "admin"},
	{ID: 2, Usuario: "socio", Contraseña: "1234", Tipo: "socio"},
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u domain.Usuario
	json.NewDecoder(r.Body).Decode(&u)

	for _, user := range usuarios {
		if user.Usuario == u.Usuario && user.Contraseña == u.Contraseña {
			token, _ := utils.GenerarToken(user.Tipo)
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}

	http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
}

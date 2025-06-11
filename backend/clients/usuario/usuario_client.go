package usuario

import (
	"backend/domain"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUserById(id int) (domain.Usuario, error) {
	var usuario domain.Usuario

	err := DB.First(&usuario, id).Error
	if err != nil {
		return domain.Usuario{}, err
	}

	return usuario, nil
}

func VerifyCredentials(email, password string) (domain.Usuario, bool) {
	var usuario domain.Usuario

	// Buscar usuario por email
	result := DB.Where("email = ?", email).First(&usuario)
	if result.Error != nil {
		log.Warn("Usuario no encontrado.")
		return domain.Usuario{}, false
	}

	// Comparar contraseña con el hash guardado
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		log.Warn("Contraseña incorrecta.")
		return domain.Usuario{}, false
	}

	log.Debug("Usuario autenticado: ", usuario.Email)
	return usuario, true
}

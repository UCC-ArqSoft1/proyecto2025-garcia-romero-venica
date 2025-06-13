package usuario

import (
	"backend/domain"
	"crypto/sha256"
	"fmt"

	log "github.com/sirupsen/logrus"
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
	hash := sha256.Sum256([]byte(password))
	passwordHash := fmt.Sprintf("%x", hash)

	result := DB.Where("email = ? AND password = ?", email, passwordHash).First(&usuario)
	if result.Error != nil {
		log.Warn("Credenciales inválidas o usuario no encontrado.")
		return domain.Usuario{}, false
	}

	log.Debug("Usuario autenticado: ", usuario.Email)
	return usuario, true
}


/*
func VerifyCredentials(email, password string) (domain.Usuario, bool) {
	var usuario domain.Usuario

	result := Db.Where("email = ?", email).First(&usuario)
	if result.Error != nil {
		log.Warn("Usuario no encontrado.")
		return domain.Usuario{}, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		log.Warn("Contraseña incorrecta.")
		return domain.Usuario{}, false
	}

	log.Debug("Usuario autenticado: ", usuario.Email)
	return usuario, true
}
*/

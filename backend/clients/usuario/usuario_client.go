package usuario

import(
	"backend/domain"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
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

	result := DB.Where("email = ? AND password = ?", email, password).First(&usuario)

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

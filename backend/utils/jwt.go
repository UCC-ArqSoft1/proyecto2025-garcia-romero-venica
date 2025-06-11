package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("mi_clave_secreta") // poné esto en variable de entorno si querés mejorar seguridad

func GenerarJWT(userID int, rol string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"rol":    rol,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidarJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

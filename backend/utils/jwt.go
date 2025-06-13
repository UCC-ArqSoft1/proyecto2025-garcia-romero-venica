package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("clave-secreta-super-segura")

// Claims es la estructura que se guarda dentro del token
type Claims struct {
	UserID int `json:"user_id"`
	Rol    int `json:"rol"`
	jwt.RegisteredClaims
}

// GenerarJWT crea un token firmado con el rol y user ID
func GenerarJWT(userID int, rol int) (string, error) {
	expiracion := time.Now().Add(2 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Rol:    rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiracion),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidarJWT analiza y valida un token existente
func ValidarJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
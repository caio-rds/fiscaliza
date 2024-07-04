package services

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte("$2a$10$S9YUva5WArXoAcP0zHNM6uQRxAhWpj61ub6TqtyDDHWg5tYqPEeEu")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJwt(username string) (string, error) {
	expirationTime := jwt.TimeFunc().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

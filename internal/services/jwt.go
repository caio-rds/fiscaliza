package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte("$2a$10$S9YUva5WArXoAcP0zHNM6uQRxAhWpj61ub6TqtyDDHWg5tYqPEeEu")
var ErrTokenExpired = errors.New("token is expired")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJwt(username string) (string, error) {
	expirationTime := jwt.TimeFunc().Add(2 * time.Hour)
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

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return claims, ErrTokenExpired
			}
		}
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.Username, nil
	} else {
		return "", err
	}
}

func GenerateJWT(username string) (string, error) {
	claims := MyCustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "books-backend",
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

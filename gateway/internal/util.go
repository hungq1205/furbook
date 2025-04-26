package internal

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("as you have seen, a very secret key")
var salt = "salty"

func Hash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error while hashing password: %v", err)
	}
	return string(hashedPassword)
}

func GenerateJwt(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString(secretKey)
}

func ParseJwt(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return username, nil
}

package internal

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("as you have seen, a very secret key")

func GenerateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatalf("Error while generating salt: %v", err)
	}
	return base64.StdEncoding.EncodeToString(salt)
}

func Hash(password, salt string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error while hashing password: %v", err)
	}
	return string(hashedPassword)
}

func CompareHashAndPassword(hashedPassword, password, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
	return err == nil
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

package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken() (string, error) {
	claims := jwt.MapClaims{}

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET_KEY")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

package service

import (
	"blog/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	SecretKey string
	Issuer    string
}

func NewJwtService() JwtService {
	return JwtService{
		SecretKey: GetSecretKey(),
		Issuer:    "secret",
	}
}
func GetSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "secret"
	}
	return secretKey

}
func (j *JwtService) GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})
	fmt.Println("token ", token)
	tokenString, err := token.SignedString(j.SecretKey)
	return tokenString, err
}

func (j *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signinig method %v", t.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})
}

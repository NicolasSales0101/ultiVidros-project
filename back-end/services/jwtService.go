package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	jwtIssure := os.Getenv("JWT_ISSURE")

	return &jwtService{
		secretKey: jwtSecret,
		issure:    jwtIssure,
	}
}

type Claim struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id string) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

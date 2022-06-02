package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	secretKey = []byte("sample_key")
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClams struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
}

func NewJWTService() JWTService {
	return &jwtService{}
}

func (service *jwtService) GenerateToken(userID string) (string, error) {
	clams := &jwtCustomClams{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer: "ryanzhou",
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	return token.SignedString(secretKey)
}

func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}
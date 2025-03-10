package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/go_web/internal/models"
	"github.com/go_web/internal/repository"
	"github.com/golang-jwt/jwt"
	"log"
	"strconv"
	"time"
)

const (
	signKey = "dewo3032NC0#01mdvfd,mpPp4qm4pcwefrr"
	salt    = "sekmcoemsp"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) CreateUser(client models.User) error {
	client.Password = r.generatePasswordHash(client.Password)
	return r.repo.CreateUser(client)
}

func (r *AuthService) GenerateJWTToken(username string, password string) (string, error) {
	id, err := r.repo.GetUserID(username, r.generatePasswordHash(password))
	if id == 0 || err != nil {
		log.Println(r.generatePasswordHash(password))
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":      id,
		"timeNow": time.Now().Unix(),
		"timer":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenSignature, err := token.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}

	return tokenSignature, nil
}

func (r *AuthService) ParsingJWTToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	idValue, exists := claims["ID"]
	if !exists {
		return "", fmt.Errorf("ID not found in token")
	}

	if idStr, ok := idValue.(string); ok {
		return idStr, nil
	}

	if idFloat, ok := idValue.(float64); ok {
		return fmt.Sprintf("%.0f", idFloat), nil
	}

	return "", err
}

func (r *AuthService) GetID(username string, password string) (string, error) {
	id, err := r.repo.GetUserID(username, r.generatePasswordHash(password))
	if id == 0 || err != nil {
		return "", err
	}
	idString := strconv.Itoa(id)
	return idString, nil
}

func (r *AuthService) generatePasswordHash(password string) string {
	passwordHash := sha1.New()
	passwordHash.Write([]byte(password))

	return hex.EncodeToString(passwordHash.Sum([]byte(salt)))
}

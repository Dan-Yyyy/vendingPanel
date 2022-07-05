package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/message"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

type AuthService struct {
	r repository.Authorisation
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(r repository.Authorisation) *AuthService {
	return &AuthService{r: r}
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New(message.UnknownTokenClaims)
	}

	return claims.UserId, nil
}

func (s *AuthService) GetUser(email string, password string) (*models.User, error) {
	user, err := s.r.GetUser(email, generatePasswordHash(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) GenerateToken(user models.User) (string, error) {
	tokenTTL, err := strconv.Atoi(os.Getenv("JWT_TOKEN_TTL"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tokenTTL) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.r.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *AuthService) GetUserById(userId int) (models.User, error) {
	return s.r.GetUserById(userId)
}

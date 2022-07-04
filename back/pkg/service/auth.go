package service

import (
	"crypto/sha1"
	"fmt"
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

var (
	singingKey = "fsdfda8s23fdzf323rfds"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s AuthService) GenerateToken(email string, password string) (string, error) {
	user, err := s.r.GetUser(email, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

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

func NewAuthService(r repository.Authorisation) *AuthService {
	return &AuthService{r: r}
}

func (s AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.r.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

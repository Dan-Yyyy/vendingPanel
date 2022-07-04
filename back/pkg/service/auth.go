package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/models"
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

type AuthService struct {
	r repository.Authorisation
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

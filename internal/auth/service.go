package auth

import (
	"errors"

	"github.com/sandu9618/food-ordering-backend/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepo *user.Repository
}

func (s *Service) Register(newUser *user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newUser.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)

	return s.UserRepo.Create(newUser)
}

func (s *Service) Login(
	email string,
	password string,
) (string, error) {
	foundUser, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return GenerateToken(foundUser.ID, foundUser.TenantID, foundUser.Role)
}

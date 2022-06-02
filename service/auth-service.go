package service

import (
	"fmt"
	"ginjwt/repo"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Verify(email, password string) (bool, error)
}

type authService struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Verify(email string, password string) (bool, error) {
	user, err := s.userRepo.FindUserByEmail(email)
	if err != nil {
		return false, fmt.Errorf("find user by email failed: %v", err)
	}
	if user == nil || user.ID == 0 {
		return false, nil
	}

	return comparePassword(user.Password, password), nil
}

func comparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
package service

import (
	"errors"
	"ginjwt/dto"
	"ginjwt/entity"
	"ginjwt/repo"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists = errors.New("user alreay exists")
	ErrUserNotExists = errors.New("user not exists")
)

type UserService interface {
	FindUserByID(userID string) (*dto.UserResponse, error)
	FindUserByEmail(email string) (*dto.UserResponse, error)
	CreateUser(req *dto.RegisterRequest) (*dto.UserResponse, error)
	UpdateUser(req *dto.UpdateUserRequest) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) FindUserByID(userID string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	return newUserResponse(user), nil
}

func (s *userService) FindUserByEmail(userEmail string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}
	return newUserResponse(user), nil
}

func (s *userService) CreateUser(req *dto.RegisterRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user = &entity.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(hashedPassword),
	}
	log.Printf("user: %+v\n", user)
	if err = s.userRepo.InsertUser(user); err != nil {
		return nil, err
	}
	return newUserResponse(user), nil
}

func (s *userService) UpdateUser(req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindUserByID(strconv.Itoa(int(req.ID)))
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID == 0 {
		return nil, ErrUserNotExists
	}

	user.Name = req.Name
	user.Email = req.Email
	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}
	return newUserResponse(user), nil
}

func newUserResponse(user *entity.User) *dto.UserResponse {
	if user == nil || user.ID == 0 {
		return nil
	}

	return &dto.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
}
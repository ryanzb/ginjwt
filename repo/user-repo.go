package repo

import (
	"ginjwt/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindUserByID(id string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	InsertUser(user *entity.User) error
	UpdateUser(user *entity.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) FindUserByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *userRepo) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (r *userRepo) InsertUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}
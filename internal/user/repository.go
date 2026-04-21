package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	DeleteUser(id string) error
	UpdateUser(id string, user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetUserByID(id string) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *repository) UpdateUser(id string, user *User) error {
	return r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

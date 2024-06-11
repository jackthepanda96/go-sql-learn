package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	Email     string
	Phone     string
	BirthDate time.Time `gorm:"type:date"`
	Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) *UserModel {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Login(email string, password string) (User, error) {
	var result User
	err := um.db.Where("email = ? AND password = ?", email, password).First(&result).Error
	if err != nil {
		return User{}, err
	}
	return result, nil
}

func (um *UserModel) Register(newUser User) (bool, error) {
	err := um.db.Create(&newUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

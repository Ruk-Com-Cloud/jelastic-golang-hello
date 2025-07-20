package infrastructure

import (
	"jelastic-golang-hello/internal/domain"
)

type UserModel struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex;not null"`
}

func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ToDomain() *domain.User {
	return &domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func FromDomain(user *domain.User) *UserModel {
	return &UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
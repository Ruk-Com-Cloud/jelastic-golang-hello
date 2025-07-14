package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserData   = errors.New("invalid user data")
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, ErrInvalidUserData
	}
	if email == "" {
		return nil, ErrInvalidUserData
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}

func (u *User) UpdateName(name string) error {
	if name == "" {
		return ErrInvalidUserData
	}
	u.Name = name
	return nil
}

func (u *User) UpdateEmail(email string) error {
	if email == "" {
		return ErrInvalidUserData
	}
	u.Email = email
	return nil
}
package application

import (
	"context"

	"jelastic-golang-hello/internal/domain"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return s.userRepo.GetAll(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, name, email string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		if err := user.UpdateName(name); err != nil {
			return nil, err
		}
	}

	if email != "" {
		if err := user.UpdateEmail(email); err != nil {
			return nil, err
		}
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(ctx, id)
}
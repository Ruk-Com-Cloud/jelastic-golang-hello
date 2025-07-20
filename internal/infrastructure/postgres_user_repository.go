package infrastructure

import (
	"context"
	"errors"

	"jelastic-golang-hello/internal/domain"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	userModel := FromDomain(user)
	
	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return err
	}
	
	user.ID = userModel.ID
	return nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var userModel UserModel
	
	if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	
	return userModel.ToDomain(), nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var userModel UserModel
	
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	
	return userModel.ToDomain(), nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	var userModels []UserModel
	
	if err := r.db.WithContext(ctx).Find(&userModels).Error; err != nil {
		return nil, err
	}
	
	users := make([]*domain.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = userModel.ToDomain()
	}
	
	return users, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	userModel := FromDomain(user)
	
	result := r.db.WithContext(ctx).Save(userModel)
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	
	return nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&UserModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	
	return nil
}
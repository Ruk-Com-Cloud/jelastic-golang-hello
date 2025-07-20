package seeder

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"jelastic-golang-hello/internal/infrastructure"
)

type UserSeeder struct{}

func NewUserSeeder() *UserSeeder {
	return &UserSeeder{}
}

func (us *UserSeeder) GetName() string {
	return "UserSeeder"
}

func (us *UserSeeder) Run(ctx context.Context, db *gorm.DB) error {
	users := []infrastructure.UserModel{
		{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		{
			Name:  "Jane Smith",
			Email: "jane.smith@example.com",
		},
		{
			Name:  "Bob Johnson",
			Email: "bob.johnson@example.com",
		},
		{
			Name:  "Alice Brown",
			Email: "alice.brown@example.com",
		},
		{
			Name:  "Charlie Wilson",
			Email: "charlie.wilson@example.com",
		},
		{
			Name:  "Diana Martinez",
			Email: "diana.martinez@example.com",
		},
		{
			Name:  "Eve Davis",
			Email: "eve.davis@example.com",
		},
		{
			Name:  "Frank Miller",
			Email: "frank.miller@example.com",
		},
		{
			Name:  "Grace Lee",
			Email: "grace.lee@example.com",
		},
		{
			Name:  "Henry Taylor",
			Email: "henry.taylor@example.com",
		},
	}

	for _, user := range users {
		// Check if user already exists
		var existingUser infrastructure.UserModel
		result := db.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser)
		
		if result.Error == nil {
			// User exists, skip
			fmt.Printf("User %s already exists, skipping\n", user.Email)
			continue
		}
		
		if result.Error != gorm.ErrRecordNotFound {
			// Real error occurred
			return fmt.Errorf("error checking existing user %s: %w", user.Email, result.Error)
		}
		
		// User doesn't exist, create it
		if err := db.WithContext(ctx).Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user %s: %w", user.Email, err)
		}
		
		fmt.Printf("Created user: %s (%s)\n", user.Name, user.Email)
	}

	return nil
}

func (us *UserSeeder) Rollback(ctx context.Context, db *gorm.DB) error {
	emails := []string{
		"john.doe@example.com",
		"jane.smith@example.com",
		"bob.johnson@example.com",
		"alice.brown@example.com",
		"charlie.wilson@example.com",
		"diana.martinez@example.com",
		"eve.davis@example.com",
		"frank.miller@example.com",
		"grace.lee@example.com",
		"henry.taylor@example.com",
	}

	for _, email := range emails {
		result := db.WithContext(ctx).Where("email = ?", email).Delete(&infrastructure.UserModel{})
		if result.Error != nil {
			return fmt.Errorf("failed to delete user %s: %w", email, result.Error)
		}
		
		if result.RowsAffected > 0 {
			fmt.Printf("Deleted user: %s\n", email)
		}
	}

	return nil
}
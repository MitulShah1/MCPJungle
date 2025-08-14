// Package user provides user management functionality.
package user

import (
	"errors"
	"fmt"

	"github.com/mcpjungle/mcpjungle/internal"
	"github.com/mcpjungle/mcpjungle/internal/model"
	"gorm.io/gorm"
)

// UserService provides methods to manage users in the MCPJungle system.
type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateAdminUser creates an admin user in the MCPJungle system.
func (u *UserService) CreateAdminUser() (*model.User, error) {
	token, err := internal.GenerateAccessToken()
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username:    "admin",
		Role:        model.UserRoleAdmin,
		AccessToken: token,
	}
	if err := u.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create admin user: %w", err)
	}

	return &user, nil
}

// VerifyAdminToken checks if the provided token belongs to an admin user
func (u *UserService) VerifyAdminToken(token string) (*model.User, error) {
	var user model.User

	err := u.db.Where("access_token = ?", token).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin user not found")
		}

		return nil, fmt.Errorf("failed to verify admin token: %w", err)
	}

	if user.Role != model.UserRoleAdmin {
		return nil, errors.New("user is not an admin")
	}

	return &user, nil
}

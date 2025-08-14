// Package migrations provides database migration functionality.
package migrations

import (
	"fmt"

	"github.com/mcpjungle/mcpjungle/internal/model"
	"gorm.io/gorm"
)

// Migrate performs the database migration for the application.
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.McpServer{})
	if err != nil {
		return fmt.Errorf("auto‑migration failed for McpServer model: %v", err)
	}

	err = db.AutoMigrate(&model.Tool{})
	if err != nil {
		return fmt.Errorf("auto‑migration failed for Tool model: %v", err)
	}

	err = db.AutoMigrate(&model.ServerConfig{})
	if err != nil {
		return fmt.Errorf("auto‑migration failed for ServerConfig model: %v", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("auto‑migration failed for User model: %v", err)
	}

	err = db.AutoMigrate(&model.McpClient{})
	if err != nil {
		return fmt.Errorf("auto‑migration failed for McpClient model: %v", err)
	}

	return nil
}

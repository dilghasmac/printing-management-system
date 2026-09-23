package database

import (
	"fmt"
	"printing-management-system/internal/domain"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.Customer{},
	)

	if err != nil {
		return fmt.Errorf("faild to run database migration %w", err)
	}

	return nil

}

package db

import (
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Task      string         `gorm:"task"`
}

func TodoMigrations() error {
	DB.AutoMigrate(&Todos{})
	return nil
}

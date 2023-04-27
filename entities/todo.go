package entities

import (
	"time"
)

type Todo struct {
	TodoID          int       `gorm:"primaryKey"`
	ActivityGroupID int       `gorm:"not null"`
	Title           string    `gorm:"not null"`
	Priority        string    `gorm:"not null"`
	IsActive        bool      `gorm:"not null"`
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
	DeletedAt       time.Time `gorm:"not null"`
}

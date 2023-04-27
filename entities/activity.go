package entities

import (
	"time"
)

type Activity struct {
	ActivityID int       `gorm:"primaryKey"`
	Title      string    `gorm:"not null"`
	Email      string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
	DeletedAt  time.Time `gorm:"not null"`
}

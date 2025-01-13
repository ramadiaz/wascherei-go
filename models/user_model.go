package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	ID              uint   `gorm:"not null;unique;primaryKey"`
	UUID            string `gorm:"not null;unique"`
	Username        string `gorm:"not null;unique;index"`
	HashedPassword  string `gorm:"not null"`
	BusinessName    string `gorm:"not null"`
	BusinessOwner   string `gorm:"not null"`
	BusinessPhone   string `gorm:"not null"`
	BusinessAddress string `gorm:"not null"`
	BusinessLogo    string `gorm:"not null"`
	Slogan          string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`
}

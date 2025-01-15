package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model

	ID       uint   `gorm:"not null;primaryKey"`
	UUID     string `gorm:"not null;unique"`
	UserUUID string `gorm:"not null"`
	Type     string `gorm:"not null"`
	Duration uint   `gorm:"not null"`
	Price    uint   `gorm:"not null"`
	Unit     string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`

	User Users `gorm:"foreignKey:UserUUID;references:UUID"`
}

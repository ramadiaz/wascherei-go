package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model

	ID            uint    `gorm:"not null;primaryKey"`
	UUID          string  `gorm:"not null;unique;index"`
	UserUUID      string  `gorm:"not null;index"`
	ProductUUID   string  `gorm:"not null;index"`
	UnitSize      float32 `gorm:"not null"`
	TotalPrice    uint    `gorm:"not null"`
	PaymentStatus string  `gorm:"not null;default:pending"`
	Customer      string  `gorm:"not null;default:Guest"`
	CustomerPhone *string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `gorm:"index"`

	User    Users    `gorm:"foreignKey:UserUUID;references:UUID"`
	Product Products `gorm:"foreignKey:ProductUUID;references:UUID"`
}

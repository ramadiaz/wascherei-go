package main

import (
	"wascherei-go/models"
	"wascherei-go/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Client{},
		&models.Users{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}

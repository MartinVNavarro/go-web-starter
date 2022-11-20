package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "host=localhost user=admin password=admin dbname=lazy port=5432 sslmode=disable timezone=America/Los_Angeles"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	err = db.AutoMigrate(&Account{})

	db.Create(&Account{ID: 1, Name: "name"})

	var account Account

	db.First(&account, 1)

	println("Found an account : " + account.Name)
}

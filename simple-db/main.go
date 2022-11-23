package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	// gorm field tag to mark ID as a PK upon insert
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	// Replace any of the paramaters accordingly
	dsn := "host=localhost user=admin password=admin dbname=lazy port=5432 sslmode=disable timezone=America/Los_Angeles"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	// Migrate schema to db
	err = db.AutoMigrate(&Account{})

	// Create a new account Record
	db.Create(&Account{ID: 1, Name: "name"})

	var account Account

	// Finds first occurance of an Account with a primary key 1
	db.First(&account, 1)

	println("Found an account : " + account.Name)
}

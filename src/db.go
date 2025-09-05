package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	ID   uint  `gorm:"primaryKey;autoIncrement;column:id"`
	Note int   `gorm:"column:note;not null"`
}

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("prideflag.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Test{})
}

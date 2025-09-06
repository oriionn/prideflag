package database

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, context.Context) {
	db, err := gorm.Open(sqlite.Open("prideflag.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()
	db.AutoMigrate(&Test{})
	db.AutoMigrate(&Images{})
	db.AutoMigrate(&Choices{})

	return db, ctx
}

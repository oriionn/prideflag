package database

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Note   int    `gorm:"column:note;not null;default:0"`
	Total  int    `gorm:"column:total;not null;default:0"`
}

type Choices struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement;column:id"`
	TestID   uint   `gorm:"column:test_id;not null"`
	TrueFlag bool   `gorm:"column:true_flag;not null"`
	Name     string `gorm:"column:name;not null"`
	File     string `gorm:"column:file;not null"`
}

type Images struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement;column:id"`
	File string `gorm:"column:file;not null"`
}

type FlagInTheTest struct {
	gorm.Model
	TestID uint `gorm:"column:test_id;not null"`
	Flag string `gorm:"colunm:flag;not null"`
}

package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:""not null"`
}

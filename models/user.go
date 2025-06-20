package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

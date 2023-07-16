package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Username  string `gorm:"not null; unique" json:"username"`
	Password  string `gorm:"not null" json:"password"`
}

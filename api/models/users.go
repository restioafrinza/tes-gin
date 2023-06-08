package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

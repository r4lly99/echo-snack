package model

import (
	"github.com/jinzhu/gorm"
)

type Snack struct {
	gorm.Model
	Slug        string `gorm:"unique_index;not null"`
	Type        string `gorm:"not null"`
	Description string
	Price       float32 `sql:"type:decimal(10,2);"`
	Favorites   []User  `gorm:"many2many:favorites;"`
}

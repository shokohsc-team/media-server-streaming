package models

import (
	"gorm.io/gorm"
)

type Video struct {
    gorm.Model
    Path  string
	CategoryID int
    Category Category
}

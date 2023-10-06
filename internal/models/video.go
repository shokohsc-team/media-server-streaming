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

func (v *Video) Extensions() string {
    return v.Path
}

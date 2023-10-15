package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID        		uint64       `gorm:"primary_key; unique; type:uint64;`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 	`gorm:"index" json:",omitempty"`
	Path          	string     		`json:"path"`
	LibraryID 		uint64       		`gorm:"type:uint64" json:"library"`
	CollectionID 	uint64       		`gorm:"type:uint64" json:"collection"`
}
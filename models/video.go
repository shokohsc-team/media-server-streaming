package models

import (
	"time"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	ID        		int       		`gorm:"primaryKey" json:"id"`
	CreatedAt 		time.Time 		`json:"created_at"`
	UpdatedAt 		time.Time 		`json:"updated_at"`
	DeletedAt 		gorm.DeletedAt 	`gorm:"index" json:"deleted_at,omitempty"`
	Path          	string     		`json:"path"`
	LibraryID 		int       		`gorm:"type:int" json:"library"`
	CollectionID 	int       		`gorm:"type:int" json:"collection"`
}
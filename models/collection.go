package models

import (
	"time"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	ID        	int       		`gorm:"primaryKey" json:"id"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted_at,omitempty"`
	Name        string     		`json:"name"`
	Videos      []Video    		`gorm:"foreignKey:CollectionID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"videos"`
}
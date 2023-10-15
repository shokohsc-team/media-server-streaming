package models

import (
	"time"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	ID        		uint64       `gorm:"primary_key; unique; type:uint64;`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:",omitempty"`
	Name        string     		`json:"name"`
	Videos      []Video    		`gorm:"foreignKey:CollectionID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}
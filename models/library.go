package models

import (
	"time"
	"gorm.io/gorm"
)

type LibraryType string

const (
    MOVIES  LibraryType = "movies"
    TVSHOWS LibraryType = "tvshows"
)

type Library struct {
	gorm.Model
	ID        		uint64       `gorm:"primary_key; unique; type:uint64;`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:",omitempty"`
	Path        string     		`json:"path"`
	LibraryType LibraryType     `gorm:"column:type" json:"type"`
	Videos    	[]Video    		`gorm:"foreignKey:LibraryID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}
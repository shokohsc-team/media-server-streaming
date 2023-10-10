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
	ID        	int       		`gorm:"primaryKey" json:"id"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index" json:"deleted_at,omitempty"`
	Name      	string     		`json:"name"`
	Path        string     		`json:"path"`
	LibraryType LibraryType     `gorm:"column:type" json:"type"`
	Videos    	[]Video    		`gorm:"foreignKey:LibraryID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"videos"`
}
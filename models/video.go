package models

import (
    "log"
    "net/http"
    "os"
    "path/filepath"

	"gorm.io/gorm"
)

type VideoModel struct {
	DB *gorm.DB
}

type Video struct {
    gorm.Model
    Path  string
	CategoryID int
    Category Category
}

func (v *Video) MIMEType() string {
    filebytes, err := os.ReadFile(v.Path)
    if err != nil {
        log.Fatal(err)
    }

    return http.DetectContentType(filebytes)
}

func (v *Video) Filename() string {
    return filepath.Base(v.Path)
}
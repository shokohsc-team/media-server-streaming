package models

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryModel struct {
	DB *gorm.DB
}

type Category struct {
	gorm.Model
    Name  string `gorm:"index" json:"name" `
}

var categories []Category

// Use a method on the custom CategoryModel type to run the SQL query.
func (m CategoryModel) All() ([]Category) {
	m.DB.Find(&categories)

	return categories
}

func (m CategoryModel) Create(c *fiber.Ctx) (result Category, error) {
	category := new(Category)
	if err := c.BodyParser(cat); err != nil {
		return err
	}

	result := m.DB.Create(&category)

	return result
}
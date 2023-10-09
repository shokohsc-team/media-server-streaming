package models

import (
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


// Use a method on the custom CategoryModel type to run the SQL query.
func (m CategoryModel) All() ([]Category) {
	var categories []Category
	m.DB.Find(&categories)

	return categories
}

func (m CategoryModel) Create(c *fiber.Ctx) (*gorm.DB, error) {
	category := new(Category)
	if err := c.BodyParser(category); err != nil {
		return nil, err
	}
	result := m.DB.Create(&category)

	return result, nil
}
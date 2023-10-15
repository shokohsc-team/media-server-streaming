package handlers

import (
	"strconv"

	"netflix/database"
	"netflix/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateLibrary(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Library)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	newLibrary := models.Library{
		Path:      		json.Path,
		LibraryType: 	json.LibraryType,
	}
	err := db.Create(&newLibrary).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
func GetLibraries(c *fiber.Ctx) error {
	db := database.DB
	Libraries := []models.Library{}
	db.Model(&models.Library{}).Order("ID asc").Limit(100).Find(&Libraries)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    Libraries,
	})
}
func GetLibraryById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}
	library := models.Library{}
	query := models.Library{ID: id}
	err = db.First(&library, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Library not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(library)
}

func UpdateLibrary(c *fiber.Ctx) error {
	type UpdateLibraryRequest struct {
		Path      	string `json:"path"`
		LibraryType models.LibraryType `json:"type"`
	}
	db := database.DB
	json := new(UpdateLibraryRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := models.Library{}
	query := models.Library{
		ID: id,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Library not found",
		})
	}
	if json.Path != "" {
		found.Path = json.Path
	}
	if json.LibraryType != "" {
		found.LibraryType = json.LibraryType
	}
	db.Save(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
func DeleteLibrary(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := models.Library{}
	query := models.Library{
		ID: id,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Library not found",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

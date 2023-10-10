package handlers

import (
	"strconv"

	"netflix/database"
	"netflix/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCollection(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Collection)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	newCollection := models.Collection{
		Name:      json.Name,
	}
	err := db.Create(&newCollection).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
func GetCollections(c *fiber.Ctx) error {
	db := database.DB
	Collections := []models.Collection{}
	db.Model(&models.Collection{}).Order("ID asc").Limit(100).Find(&Collections)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    Collections,
	})
}
func GetCollectionById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}
	collection := models.Collection{}
	query := models.Collection{ID: id}
	err = db.First(&collection, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Collection not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(collection)
}

func UpdateCollection(c *fiber.Ctx) error {
	type UpdateCollectionRequest struct {
		Name      string `json:"name"`
	}
	db := database.DB
	json := new(UpdateCollectionRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := models.Collection{}
	query := models.Collection{
		ID: id,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Collection not found",
		})
	}
	if json.Name != "" {
		found.Name = json.Name
	}
	db.Save(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
func DeleteCollection(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := models.Collection{}
	query := models.Collection{
		ID: id,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Collection not found",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

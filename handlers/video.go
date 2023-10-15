package handlers

import (
	"strconv"

	"netflix/database"
	"netflix/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateVideo(c *fiber.Ctx) error {
	db := database.DB
	json := new(models.Video)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	library := c.Locals("library").(models.Library)
	collection := c.Locals("collection").(models.Collection)
	newVideo := models.Video{
		LibraryID: library.ID,
		CollectionID: collection.ID,
		Path:      json.Path,
	}
	err := db.Create(&newVideo).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
func GetVideos(c *fiber.Ctx) error {
	db := database.DB
	Videos := []models.Video{}
	db.Model(&models.Video{}).Order("ID asc").Limit(100).Find(&Videos)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    Videos,
	})
}
func GetVideoById(c *fiber.Ctx) error {
	db := database.DB
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}
	video := models.Video{}
	query := models.Video{ID: id}
	err = db.First(&video, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Video not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(video)
}

func UpdateVideo(c *fiber.Ctx) error {
	type UpdateVideoRequest struct {
		Path      string `json:"name"`
	}
	db := database.DB
	library := c.Locals("library").(models.Library)
	collection := c.Locals("collection").(models.Collection)
	json := new(UpdateVideoRequest)
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
	found := models.Video{}
	query := models.Video{
		ID: id,
		LibraryID: library.ID,
		CollectionID: collection.ID,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Video not found",
		})
	}
	if json.Path != "" {
		found.Path = json.Path
	}
	db.Save(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
func DeleteVideo(c *fiber.Ctx) error {
	db := database.DB
	library := c.Locals("library").(models.Library)
	collection := c.Locals("collection").(models.Collection)
	param := c.Params("id")
	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}
	found := models.Video{}
	query := models.Video{
		ID: id,
		LibraryID: library.ID,
		CollectionID: collection.ID,
	}
	err = db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Video not found",
		})
	}
	db.Delete(&found)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

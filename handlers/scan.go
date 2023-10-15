package handlers

import (
	"os"
	"fmt"
	"log"
    "path/filepath"

	"netflix/database"
	"netflix/models"
	"netflix/queue"

	"github.com/gofiber/fiber/v2"
)

type message struct {
	path string
	library uint64
}



func scan(l models.Library) error {
	var files []string

	err := filepath.Walk(l.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && (filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".mp4") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		v := models.Video{
			Path:      files[i],
			LibraryID: l.ID,
		}
		queue.Publish(v)
		if err != nil {
			return err
		}
    }

	return err
}

func ExecuteScan(c *fiber.Ctx) error {
	db := database.DB
	Libraries := []models.Library{}
	db.Model(&models.Library{}).Order("ID asc").Find(&Libraries)

	for _, library := range Libraries {
		go scan(library)
    }

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

package main

import (
	"fmt"
	"log"
	"os"
    "path/filepath"

    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"netflix/models"
)

const mediaPath = "/mnt"
var db *gorm.DB

type Env struct {
	categories models.CategoryModel
	videos models.VideoModel
}

func initDb() {
	dsn := "host=postgres user=netflix password=netflix dbname=netflix port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        log.Fatal(err)
    }

	// Migrate the schema
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Video{})
}

func scan(directory string) error {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
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

	for _, file := range files {
		fmt.Println(file)
	}

	return err
}

func main() {
	initDb()
    // Start a new fiber app
    app := fiber.New(fiber.Config{
		AppName: "netflix",
		EnablePrintRoutes: true,
		EnableSplittingOnParsers: true,
    })

	app.Use(etag.New())
	app.Use(cors.New())

	// Send a string back for GET calls to the endpoint "/"
    app.Get("/ready", func(c *fiber.Ctx) error {
		return c.Send(nil)
    }).Name("ready")

	app.Post("/category", func(c *fiber.Ctx) error {
		env := &Env{
			categories: models.CategoryModel{DB: db},
		}
		category, err := env.categories.Create(c)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(fiber.Map{
			"name": category.Name,
		})
	}).Name("postCategory")

	// app.Post("/video", func(c *fiber.Ctx) error {
	// 	env := &Env{
	// 		videos: models.VideoModel{DB: db},
	// 	}
	// 	env.videos.Create(c)
    //
	// return c.Send(nil)
	// }).Name("videoCategory")

	app.Post("/scan", func(c *fiber.Ctx) error {
		// Look for video to save to database
		go scan(mediaPath)
		return c.JSON(fiber.Map{
			"status": "scan started.",
		})
    }).Name("scan")

    app.Get("/start/:videoId", func(c *fiber.Ctx) error {
        videoId := c.Params("videoId")
		// Get info from database with videoId
		// Start ffmpeg pod, service & ingress
		// Redirect to ffmpeg ingress
		return c.RedirectToRoute("stream", fiber.Map{
			"videoId": videoId,
		})
    }).Name("start")

    app.Get("/stop/:videoId", func(c *fiber.Ctx) error {
        videoId := c.Params("videoId")
		// Get info from database with videoId
		// Delete ffmpeg ingress, service & pod
		return c.JSON(fiber.Map{
			"videoId": videoId,
		})
    }).Name("stop")

    app.Get("/stream/:videoId", func(c *fiber.Ctx) error {
        videoId := c.Params("videoId")
		// Get info from database with videoId
		// Start ffmpeg stream command
		return c.JSON(fiber.Map{
			"videoId": videoId,
		})
    }).Name("stream")

    // Listen on PORT 8080
    app.Listen(":8080")
}

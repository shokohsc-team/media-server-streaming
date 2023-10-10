package main

import (
	"fmt"
	"log"
	"os"
    "path/filepath"

    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"netflix/database"
	"netflix/router"
)

const mediaPath = "/mnt"

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
	database.ConnectDB()
    // Start a new fiber app
    app := fiber.New(fiber.Config{
		AppName: "netflix",
		EnablePrintRoutes: true,
		EnableSplittingOnParsers: true,
    })

	app.Use(etag.New())
	app.Use(cors.New())

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

	app.Post("/scan", func(c *fiber.Ctx) error {
		// Look for video to save to database
		go scan(mediaPath)
		return c.JSON(fiber.Map{
			"status": "scan started.",
		})
    }).Name("scan")

	router.Initalize(app)
	log.Fatal(app.Listen(":8080"))
}

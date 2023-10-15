package main

import (
	"log"

    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"netflix/database"
	"netflix/queue"
	"netflix/router"
)

func main() {
	database.ConnectDB()
	queue.ConnectAMQP()
	go queue.Consume()

    // Start a new fiber app
    app := fiber.New(fiber.Config{
		AppName: "netflix",
		// EnablePrintRoutes: true,
		EnableSplittingOnParsers: true,
    })

	app.Use(etag.New())
	app.Use(cors.New())

	router.Initalize(app)
	log.Fatal(app.Listen(":8080"))
}

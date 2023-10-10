package router

import (
	"netflix/handlers"
	"netflix/middleware"
	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App) {

	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Ready!")
	}).Name("ready")

	app.Use(middleware.Json)

	libraries := app.Group("/libraries")
	libraries.Post("/", handlers.CreateLibrary).Name("createLibrary")
	libraries.Get("/", handlers.GetLibraries).Name("getLibraries")
	libraries.Delete("/:id", handlers.DeleteLibrary).Name("removeLibrary")
	libraries.Get("/:id", handlers.GetLibraryById).Name("getLibrary")
	libraries.Put("/:id", handlers.UpdateLibrary).Name("updateLibrary")

	collections := app.Group("/collections")
	collections.Post("/", handlers.CreateCollection).Name("createCollection")
	collections.Get("/", handlers.GetCollections).Name("getCollections")
	collections.Delete("/:id", handlers.DeleteCollection).Name("removeCollection")
	collections.Get("/:id", handlers.GetCollectionById).Name("getCollection")
	collections.Put("/:id", handlers.UpdateCollection).Name("updateCollection")

	videos := app.Group("/videos")
	videos.Post("/", handlers.CreateVideo).Name("createVideo")
	videos.Get("/", handlers.GetVideos).Name("getVideos")
	videos.Delete("/:id", handlers.DeleteVideo).Name("removeVideo")
	videos.Get("/:id", handlers.GetVideoById).Name("getVideo")
	videos.Put("/:id", handlers.UpdateVideo).Name("updateVideo")

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}

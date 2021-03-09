package routers

import (
	"fmt"
	"os"

	_ "danilopeixoto.com/api/music/docs" // generated docs
	"danilopeixoto.com/api/music/handlers"
	"danilopeixoto.com/api/music/models"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes function
func SetupRoutes(app *fiber.App) {
	version := os.Getenv("API_VERSION")

	root := app.Group(fmt.Sprintf("/%s", version))

	api := root.Group("/api")

	api.Post("/", handlers.AddSong)
	api.Get("/", handlers.GetAllSongs)
	api.Get("/:id", handlers.GetSong)
	api.Put("/:id", handlers.UpdateSong)
	api.Delete("/:id", handlers.DeleteSong)

	app.Get("/docs/*", swagger.Handler)

	app.Use(func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusNotFound).JSON(
			&models.ErrorResponse{
				Message: "API route not found.",
			})
	})
}
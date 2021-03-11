package routers

import (
	"fmt"

	"danilopeixoto.com/api/music/config"
	_ "danilopeixoto.com/api/music/docs" // generated docs
	"danilopeixoto.com/api/music/handlers"
	"danilopeixoto.com/api/music/models"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes function
func SetupRoutes(app *fiber.App) {
	apiConfig := config.GetAPIConfig()
	root := app.Group(fmt.Sprintf("/%s", apiConfig.Version))

	api := root.Group("/api")

	api.Post("/", handlers.AddSong)
	api.Get("/", handlers.GetAllSongs)
	api.Get("/:id", handlers.GetSong)
	api.Put("/:id", handlers.UpdateSong)
	api.Delete("/:id", handlers.DeleteSong)

	root.Use("/docs", swagger.Handler)

	app.Use(func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusNotFound).JSON(
			&models.ErrorResponse{
				Message: "API route not found.",
			})
	})
}

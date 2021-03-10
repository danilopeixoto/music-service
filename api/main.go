package main

import (
	"errors"
	"fmt"

	"danilopeixoto.com/api/music/config"
	"danilopeixoto.com/api/music/database"
	"danilopeixoto.com/api/music/models"
	"danilopeixoto.com/api/music/routers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @title Music API
// @version v1
// @description A music web service.
// @contact.name Danilo Peixoto
// @contact.email danilopeixoto@outlook.com
// @license.name BSD-3-Clause
// @license.url https://github.com/danilopeixoto/music-service/LICENSE
// @host localhost:8080
// @BasePath /
func main() {
	config.LoadConfig()
	database.Connect()

	apiConfig := config.GetAPIConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(context *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			} else if _, ok := err.(*validator.ValidationErrors); ok {
				code = fiber.StatusBadRequest
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				code = fiber.StatusNotFound
			}

			return context.Status(code).JSON(
				&models.ErrorResponse{
					Message: err.Error(),
				})
		},
	})

	routers.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", apiConfig.Port))
}

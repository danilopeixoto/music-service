package handlers

import (
	"danilopeixoto.com/api/music/database"
	"danilopeixoto.com/api/music/models"
	"danilopeixoto.com/api/music/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AddSong handler
// @Summary Add song
// @Description Add song record.
// @Router /v1/api [post]
// @Param song body models.SongRequest true "Song request."
// @Accept json
// @Produce json
// @Success 200 {object} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func AddSong(context *fiber.Ctx) error {
	db := database.GetDatabase()

	songRequest := new(models.SongRequest)

	if err := context.BodyParser(songRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := utils.Validate(songRequest); err != nil {
		return err
	}

	song := &models.Song{
		Title:    songRequest.Title,
		Artist:   songRequest.Artist,
		Duration: songRequest.Duration,
	}

	if err := db.Create(song).Error; err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(song)
}

// GetAllSongs handler
// @Summary List or find song
// @Description List or find song by title.
// @Router /v1/api [get]
// @Param title query string false "Song title."
// @Accept json
// @Produce json
// @Success 200 {array} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func GetAllSongs(context *fiber.Ctx) error {
	db := database.GetDatabase()

	songQuery := new(models.SongQuery)

	if err := context.QueryParser(songQuery); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	songs := new([]models.Song)

	var err error = nil

	if len(songQuery.Title) > 0 {
		err = db.Find(songs, songQuery).Error
	} else {
		err = db.Find(songs).Error
	}

	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(songs)
}

// GetSong handler
// @Summary Find song
// @Description Find song by ID.
// @Router /v1/api/{id} [get]
// @Param id path string true "Song ID."
// @Accept json
// @Produce json
// @Success 200 {object} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func GetSong(context *fiber.Ctx) error {
	db := database.GetDatabase()

	id, err := uuid.Parse(context.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	song := new(models.Song)

	if err := db.First(song, id).Error; err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(song)
}

// UpdateSong handler
// @Summary Update song
// @Description Update song by ID.
// @Router /v1/api/{id} [put]
// @Param id path string true "Song ID."
// @Param song body models.SongRequest true "Song request."
// @Accept json
// @Produce json
// @Success 200 {object} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func UpdateSong(context *fiber.Ctx) error {
	db := database.GetDatabase()

	id, err := uuid.Parse(context.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	songRequest := new(models.SongRequest)

	if err := context.BodyParser(songRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := utils.Validate(songRequest); err != nil {
		return err
	}

	song := &models.Song{
		Title:    songRequest.Title,
		Artist:   songRequest.Artist,
		Duration: songRequest.Duration,
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Song{ID: id}).Updates(song).Error; err != nil {
			return err
		}

		if err := tx.First(song, id).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(song)
}

// DeleteSong handler
// @Summary Delete song
// @Description Delete song by ID.
// @Router /v1/api/{id} [delete]
// @Param id path string true "Song ID."
// @Accept json
// @Produce json
// @Success 200 {object} models.Song
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func DeleteSong(context *fiber.Ctx) error {
	db := database.GetDatabase()

	id, err := uuid.Parse(context.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	song := new(models.Song)

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(song, id).Error; err != nil {
			return err
		}

		if err := tx.Delete(song).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(song)
}

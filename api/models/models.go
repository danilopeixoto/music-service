package models

import (
	"time"

	uuid "github.com/google/uuid"
)

// ErrorResponse model
type ErrorResponse struct {
	Message string `json:"message,required"`
} // @name ErrorResponse

// SongRequest model
type SongRequest struct {
	Title    string `json:"title,required" validate:"required"`
	Artist   string `json:"artist,required" validate:"required"`
	Duration uint64 `json:"duration,required" validate:"required,min=0,number"`
} // @name SongRequest

// SongQuery model
type SongQuery struct {
	Title string `query:"title"`
}

// Song model
type Song struct {
	ID        uuid.UUID `json:"id,required" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title     string    `json:"title,required"`
	Artist    string    `json:"artist,required"`
	Duration  uint64    `json:"duration,required"`
	CreatedAt time.Time `json:"created_at,required"`
	UpdatedAt time.Time `json:"updated_at,required"`
} // @name Song

// GetDatabaseModels function
func GetDatabaseModels() []interface{} {
	return []interface{}{
		&Song{},
	}
}

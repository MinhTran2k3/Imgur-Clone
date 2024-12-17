package models

import (
	"time"
)

// Image represents an image uploaded by a user
type Image struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

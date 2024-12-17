package models

import (
	"time"
)

type Image struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

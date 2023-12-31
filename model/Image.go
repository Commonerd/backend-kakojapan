package model

import "time"

type Image struct {
	ID uint `json:"id" gorm:"primaryKey"`
	//UserID      string    `json:"user_id" gorm:"unique"`
	Title       string    `json:"title" gorm:"unique"`
	Description string    `json:"description"`
	Lat         float64   `json:"lat"`
	Lng         float64   `json:"lng"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ImageResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	//UserID string `json:"user_id" gorm:"unique"`
	Title string `json:"title" gorm:"unique"`
}

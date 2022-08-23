package entity

import (
	"time"
)

type (
	Cake struct {
		Id          int       `json:"id"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Rating      float64   `json:"rating"`
		Image       string    `json:"image" validate:"required"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	ListCake struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}
	Error struct {
		Code    int    `json:"code"`
		Status  string `json:"error"`
		Message string `json:"message"`
	}
)

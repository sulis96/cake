package entity

import (
	"time"
)

type (
	Cake struct {
		Id          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Rating      float64   `json:"rating"`
		Image       string    `json:"image"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	ListCake struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}

	Response struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	Error struct {
		Code    int    `json:"code"`
		Status  string `json:"error"`
		Message string `json:"message"`
	}
)

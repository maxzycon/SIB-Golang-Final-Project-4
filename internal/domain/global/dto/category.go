package dto

import "time"

type CategoryRow struct {
	ID        uint               `json:"id"`
	Type      string             `json:"type"`
	UpdatedAt time.Time          `json:"updated_at"`
	CreatedAt time.Time          `json:"created_at"`
	Task      []*TaskCategoryRow `json:"Task"`
}

type TaskCategoryRow struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PayloadCategory struct {
	Type string `json:"type"`
}

type PayloadUpdateCategory struct {
	Type string `json:"type"`
}

type CategoryCreteResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryUpdateResponse struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

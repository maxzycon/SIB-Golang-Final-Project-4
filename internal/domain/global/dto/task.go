package dto

import "time"

type TaskWrappper struct {
	Tasks []*TaskRow `json:"social_medias"`
}

type TaskRow struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Status      bool         `json:"status"`
	Description string       `json:"description"`
	UserID      uint         `json:"user_id"`
	CategoryID  uint         `json:"category_id"`
	CreatedAt   time.Time    `json:"created_at"`
	User        *UserTaskRow `json:"User"`
}

type UserTaskRow struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type PayloadTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}

type PayloadUpdateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PayloadUpdateStatusTask struct {
	Status bool `json:"status"`
}

type PayloadUpdateCategoryTask struct {
	CategoryID uint `json:"category_id"`
}

type TaskCreteResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskUpdateResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

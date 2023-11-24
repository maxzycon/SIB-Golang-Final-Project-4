package dto

import "time"

type ProductRow struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserProductRow struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type PayloadProduct struct {
	Title      string `json:"title"`
	Price      uint64 `json:"price"`
	Stock      uint64 `json:"stock"`
	CategoryID uint   `json:"category_Id"`
}

type PayloadUpdateProduct struct {
	Title      string `json:"title"`
	Price      uint64 `json:"price"`
	Stock      uint64 `json:"stock"`
	CategoryID uint   `json:"category_Id"`
}

type PayloadUpdateStatusProduct struct {
	Status bool `json:"status"`
}

type PayloadUpdateCategoryProduct struct {
	CategoryID uint `json:"category_id"`
}

type ProductCreteResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint      `json:"category_Id"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProductUpdateWrappper struct {
	Products *ProductUpdateResponse `json:"product"`
}

type ProductUpdateResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint      `json:"CategoryId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

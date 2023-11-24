package dto

import "time"

type CategoryRow struct {
	ID        uint                 `json:"id"`
	Type      string               `json:"type"`
	UpdatedAt time.Time            `json:"updated_at"`
	CreatedAt time.Time            `json:"created_at"`
	Products  []ProductCategoryRow `json:"products"`
}

type ProductCategoryRow struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     uint64    `json:"description"`
	Stock     uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PayloadCategory struct {
	Type string `json:"type"`
}

type PayloadUpdateCategory struct {
	Type string `json:"type"`
}

type CategoryCreteResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint64    `json:"sold_product_amoun"`
	CreatedAt         time.Time `json:"created_at"`
}

type CategoryUpdateResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint64    `json:"sold_product_amoun"`
	UpdatedAt         time.Time `json:"updated_at"`
}

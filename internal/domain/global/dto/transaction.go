package dto

import "time"

type PayloadTransaction struct {
	ProductID uint64 `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
}

type ResponseCreateTransaction struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type TransactionBill struct {
	TotalPrice   uint64 `json:"total_price"`
	Quantity     uint64 `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type MyTransaction struct {
	ID         uint                 `json:"id"`
	ProductID  uint                 `json:"product_id"`
	UserID     uint                 `json:"user_id"`
	Quantity   uint64               `json:"quantity"`
	TotalPrice uint64               `json:"total_price"`
	Product    ProductMyTransaction `json:"Product"`
}

type ProductMyTransaction struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint      `json:"category_Id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserTransaction struct {
	ID         uint                   `json:"id"`
	ProductID  uint                   `json:"product_id"`
	UserID     uint                   `json:"user_id"`
	Quantity   uint64                 `json:"quantity"`
	TotalPrice uint64                 `json:"total_price"`
	Product    ProductUserTransaction `json:"Product"`
	User       ProductUserInfo        `json:"User"`
}

type ProductUserInfo struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   uint64    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductUserTransaction struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint      `json:"category_Id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

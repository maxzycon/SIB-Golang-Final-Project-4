package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null;type:varchar(100);index:unique"`
	Email    string `gorm:"not null;type:varchar(100);index:unique"`
	Password string `gorm:"not null;type:varchar(100)"`
	Role     uint   `gorm:"not null;"`
	Balance  uint   `gorm:"not null;default:0"`
}

type Category struct {
	gorm.Model
	Type              string
	SoldProductAmount uint64
	Product           []Product
}

type Product struct {
	gorm.Model
	Title      string
	Price      uint64
	Stock      uint64
	CategoryID uint
	Category   Category
}

type TransactionHistory struct {
	gorm.Model
	ProductID  uint
	Product    Product
	UserID     uint
	User       User
	Quantity   uint64
	TotalPrice uint64
}

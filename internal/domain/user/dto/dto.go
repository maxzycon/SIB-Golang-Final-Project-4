package dto

import "time"

type UserRow struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Role     uint    `json:"role"`
	Image    *string `json:"image"`
}

type UserRowDetail struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRowDetailUpdateRes struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PayloadUpdateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type PayloadLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type PayloadUpdateProfile struct {
	Password string `json:"password"`
}

type PayloadCreateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type LoginRes struct {
	AccessToken string `json:"token"`
}

type UserPaginatedRow struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Merchant string `json:"merchant"`
	Role     uint   `json:"role"`
}

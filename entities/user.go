package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null;unique" json:"email" form:"email"`
	Phone    string `gorm:"not null;unique" json:"phone" form:"phone"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Role     string `gorm:"default:user" json:"role" form:"role" `
}

type GetUserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `gorm:"not null" json:"name" form:"name"`
	Email string `gorm:"not null;unique" json:"email" form:"email"`
	Phone string `gorm:"not null;unique" json:"phone" form:"phone"`
}

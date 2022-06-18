package entities

type User struct {
	Name     string `gorm:"not null" json:"name" form:"name"`
	Email    string `gorm:"not null;unique" json:"email" form:"email"`
	Phone    string `gorm:"not null;unique" json:"phone" form:"phone"`
	Password string `gorm:"not null" json:"password" form:"password"`
	Role     string `gorm:"default:user" json:"role" form:"role" `
}

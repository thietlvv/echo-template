package entities

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// TableName ...
func (b *User) TableName() string {
	return "user"
}

package entities

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// TableName ...
func (user *User) TableName() string {
	return "user"
}

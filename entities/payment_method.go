package entities

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	MethodName        string `gorm:"not_null" json:"method_name"`
	MethodDescription string `json:"method_description"`
	Status            int    `gorm:"not_null" json:"status"`
}

// TableName ...
func (paymentMethod *PaymentMethod) TableName() string {
	return "payment_method"
}

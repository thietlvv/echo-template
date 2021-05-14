package entities

import "gorm.io/gorm"

type BillingOrder struct {
	gorm.Model
	OrderID           string `gorm:"not_null" json:"order_id"`
	UserID            string `gorm:"not_null" json:"user_id"`
	ProductID         string `json:"product_id"`
	PromotionCode     string `json:"product_code"`
	PromotionCodeType int    `json:"product_code_type"`
	ProductPrice      int    `gorm:"not_null" json:"product_price"`
	DiscountAmount    int    `gorm:"not_null" json:"discount_amount"`
	PaymentPrice      int    `gorm:"not_null" json:"payment_price"`
	OrderDate         string `gorm:"not_null" json:"order_date"`
	PaymentProvider   string `gorm:"not_null" json:"payment_provider"`
	Note              string `json:"note"`
}

// TableName ...
func (billingOrder *BillingOrder) TableName() string {
	return "billing_order"
}

package entities

import "gorm.io/gorm"

type PromotionGiftCode struct {
	gorm.Model
	EventCode       string `gorm:"not_null" json:"event_code"`
	PromotionCode   string `gorm:"not_null" json:"promotion_code"`
	PromotionSerial string `json:"promotion_serial"`
	Status          int    `gorm:"not_null; default:10" json:"status"`
	Note            string `json:"note"`
}

// TableName ...
func (promotionGiftCode *PromotionGiftCode) TableName() string {
	return "promotion_gift_code"
}

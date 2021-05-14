package entities

import "gorm.io/gorm"

type PromotionEvent struct {
	gorm.Model
	EventCode                  string `gorm:"not_null" json:"event_code"`
	EventName                  string `gorm:"not_null" json:"event_name"`
	EventDescription           string `json:"event_description"`
	Status                     int    `gorm:"not_null; default:0" json:"status"`
	StartDate                  string `gorm:"not_null" json:"start_date"`
	EndDate                    string `gorm:"not_null" json:"end_date"`
	PromotionType              int    `gorm:"not_null; default:1" json:"promotion_type"`
	PromotionCodeType          int    `gorm:"not_null; default:1" json:"promotion_code_type"`
	PromotionProductID         string `json:"promotion_product_id"`
	CodePrefix                 string `json:"code_prefix"`
	MultipleTimesPromotionCode string `json:"multiple_times_promotion_code"`
	CodeAmount                 int    `json:"code_amount"`
	RequestDate                string `json:"request_date"`
	RequestUserID              string `json:"request_user_id"`
	ApproveDate                string `json:"approve_date"`
	ApproveUserID              string `json:"approve_user_id"`
	Note                       string `json:"note"`
}

// TableName ...
func (promotionEvent *PromotionEvent) TableName() string {
	return "promotion_event"
}

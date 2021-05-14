package entities

import "gorm.io/gorm"

type PaymentMoneyTransferRate struct {
	gorm.Model
	PaymentProviderID string  `gorm:"not_null" json:"payment_provider_id"`
	TransferRate      float64 `gorm:"not_null; type:decimal(3,2)" json:"transfer_rate"`
	RoundingType      int     `gorm:"not_null" json:"rounding_type"`
}

// TableName ...
func (paymentMoneyTransferRate *PaymentMoneyTransferRate) TableName() string {
	return "payment_money_transfer_rate"
}

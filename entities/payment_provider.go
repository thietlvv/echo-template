package entities

import "gorm.io/gorm"

type PaymentProvider struct {
	gorm.Model
	ProviderName        string `gorm:"not_null" json:"provider_name"`
	ProviderDescription string `json:"provider_description"`
	PaymentMethodID     string `json:"payment_method_id"`
	Status              int    `gorm:"not_null" json:"status"`
	PlatformAvailable   string `json:"platform_available"`
}

// TableName ...
func (paymentProvider *PaymentProvider) TableName() string {
	return "payment_method"
}

package entities

import "gorm.io/gorm"

type WalletUser struct {
	gorm.Model
	WalletID string `gorm:"not_null" json:"wallet_id"`
	UserID   string `gorm:"not_null" json:"user_id"`
	Balance  int    `gorm:"not_null" json:"balance"`
	Status   int    `gorm:"not_null" json:"status"`
	Note     string `json:"note"`
}

// TableName ...
func (walletUser *WalletUser) TableName() string {
	return "wallet_user"
}

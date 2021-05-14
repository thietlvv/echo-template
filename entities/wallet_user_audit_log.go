package entities

import "gorm.io/gorm"

type WalletUserAuditLog struct {
	gorm.Model
	SourceType  int    `gorm:"not_null" json:"source_type"`
	TxnID       string `json:"txn_id"`
	TransferID  string `json:"transfer_id"`
	UserID      string `gorm:"not_null" json:"user_id"`
	WalletID    string `gorm:"not_null" json:"wallet_id"`
	Amount      int    `gorm:"not_null" json:"amount"`
	CreditDebit string `gorm:"not_null" json:"credit_debit"`
	ChangeTime  string `gorm:"not_null" json:"change_time"`
}

// TableName ...
func (walletUserAuditLog *WalletUserAuditLog) TableName() string {
	return "wallet_user_audit_log"
}

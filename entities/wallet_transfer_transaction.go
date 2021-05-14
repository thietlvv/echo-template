package entities

import "gorm.io/gorm"

type WalletTransferTransaction struct {
	gorm.Model
	OriginTxnTransferID string `json:"origin_txn_transfer_id"`
	TxnTransferID       string `gorm:"not_null" json:"txn_transfer_id"`
	TxnType             int    `gorm:"not_null" json:"txn_type"`
	PaymentTxnID        int    `json:"payment_txn_id"`
	MakeActionUserID    string `gorm:"not_null" json:"make_action_user_id"`
	WalletID            string `gorm:"not_null" json:"wallet_id"`
	Amount              string `gorm:"not_null" json:"amount"`
	CreditDebit         string `gorm:"not_null" json:"credit_debit"`
	Status              string `gorm:"not_null" json:"status"`
	ErrorCode           string `gorm:"not_null" json:"error_code"`
	ErrorString         string `gorm:"not_null" json:"error_string"`
	ExchangeTime        string `gorm:"not_null" json:"exchange_time"`
	Note                string `json:"note"`
}

// TableName ...
func (walletTransferTransaction *WalletTransferTransaction) TableName() string {
	return "wallet_transfer_transaction"
}

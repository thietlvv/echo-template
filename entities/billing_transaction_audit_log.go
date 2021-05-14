package entities

import "gorm.io/gorm"

type BillingTransactionAuditLog struct {
	gorm.Model
	TxnID           string `gorm:"not_null" json:"txn_id"`
	PartnerTxnID    string `gorm:"not_null" json:"partner_txn_id"`
	PartnerSubTxnID string `gorm:"not_null" json:"partner_sub_txn_id"`
	OrderID         string `gorm:"not_null" json:"order_id"`
	InvoiceID       string `gorm:"not_null" json:"invoice_id"`
	Status          int    `gorm:"not_null" json:"status"`
	ErrorCode       string `json:"error_code"`
	ErrorString     string `json:"error_string"`
}

// TableName ...
func (billingTransactionAuditLog *BillingTransactionAuditLog) TableName() string {
	return "billing_transaction_audit_log"
}
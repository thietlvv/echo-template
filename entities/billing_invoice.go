package entities

import "gorm.io/gorm"

type BillingInvoice struct {
	gorm.Model
	InvoiceID          string `gorm:"not_null" json:"invoice_id"`
	TxnID              string `gorm:"not_null" json:"txn_id"`
	OrderID            string `json:"order_id"`
	ProductID          string `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductType        int    `json:"product_type"`
	ProductPrice       int    `json:"product_price"`
	UserID             int    `json:"user_id"`
	UserName           int    `json:"user_name"`
	PromotionCode      string `json:"promotion_code"`
	PromotionCodeType  int    `json:"promotion_code_type"`
	DiscountAmount     int    `json:"discount_amount"`
	PaymentPrice       int    `json:"payment_price"`
	OrderDate          string `json:"order_date"`
	PaymentProvider    string `json:"payment_provider"`
	PaymentSuccessDate string `json:"payment_success_date"`
	PointExchange      int    `json:"point_exchange"`
}

// TableName ...
func (billingInvoice *BillingInvoice) TableName() string {
	return "billing_invoice"
}

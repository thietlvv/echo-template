package migrations

import (
	"billing/db/main_db"
	"billing/entities"
)

// MigrateDataTable ...
func MigrateDataTable() {
	main_db.DB.AutoMigrate(
		&entities.User{},
		&entities.SellProduct{},
		&entities.BillingInvoice{},
		&entities.BillingOrder{},
		&entities.BillingPaymentTransaction{},
		&entities.BillingTransactionAuditLog{},
		&entities.PaymentMethod{},
		&entities.PaymentMoneyTransferRate{},
		&entities.PaymentProvider{},
		&entities.PromotionEvent{},
		&entities.PromotionGiftCode{},
		&entities.WalletTransferTransaction{},
		&entities.WalletUserAuditLog{},
		&entities.WalletUser{},
	)
}

package entities

import "gorm.io/gorm"

type SellProduct struct {
	gorm.Model
	ProductID          string `gorm:"not_null" json:"product_id"`
	ProductName        string `gorm:"not_null" json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductType        int    `gorm:"not_null; default:1" json:"product_type"`
	Status             int    `gorm:"not_null; default:0" json:"status"`
	StartSellingDate   string `gorm:"not_null" json:"start_selling_date"`
	EndSellingDate     string `gorm:"not_null" json:"end_selling_date"`
	SellingOnPlatform  string `json:"selling_on_platform"`
	Price              int    `gorm:"not_null" json:"price"`
	IosProductID       string `json:"ios_product_id"`
	AndroidProductID   string `json:"android_product_id"`
}

// TableName ...
func (sellProduct *SellProduct) TableName() string {
	return "sell_product"
}

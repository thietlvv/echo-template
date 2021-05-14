package models

import (
	mainDb "billing/db/main_db"
	. "billing/entities"
	"log"

	"fmt"
)

// SellProduct ...
type SellProductModel struct {
}

//GetAllFetch all sell product data
func (model *SellProductModel) GetAll() (sellProduct []SellProduct, err error) {
	if err := mainDb.DB.Find(&sellProduct).Error; err != nil {
		return sellProduct, err
	}
	return sellProduct, nil
}

//CreateSellProduct ... Insert New data
func (model *SellProductModel) CreateSellProduct(sellProduct *SellProduct) (err error) {
	log.Printf("[CreateSellProduct] Received message in user service: %+v\n", sellProduct)
	if err = mainDb.DB.Create(sellProduct).Error; err != nil {
		return err
	}
	return nil
}

//GetSellProductByID ... Fetch only one sell product by Id
func (model *SellProductModel) GetSellProductByID(sellProduct *SellProduct, id string) (err error) {
	if err = mainDb.DB.Where("id = ?", id).First(sellProduct).Error; err != nil {
		return err
	}
	return nil
}

//UpdateSellProduct ... Update sell product
func (model *SellProductModel) UpdateSellProduct(sellProduct *SellProduct, id string) (err error) {
	fmt.Println(sellProduct)
	mainDb.DB.Save(sellProduct)
	return nil
}

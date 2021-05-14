package sell_product

import (
	"net/http"

	"billing/entities"
	"billing/helper"
	"billing/models"
	"billing/utils/response"

	"github.com/gin-gonic/gin"
)

//GetSellProduct ... Get all sell product
func GetSellProduct(c *gin.Context) {
	var productModel models.SellProductModel
	requestID := c.GetString("x-request-id")
	helper.Logger(requestID, "").Infoln("RequestID= ", requestID)

	var products []entities.SellProduct
	products, err := productModel.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response.Success(c, "", products)
	}
}

package sell_product

import (
	"github.com/gin-gonic/gin"
)

func Routes(g *gin.RouterGroup) {
	grpUser := g.Group("/products")

	grpUser.GET("", GetSellProduct)

}

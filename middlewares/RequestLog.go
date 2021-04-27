package middleware

import (
	l "billing/utils/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	"moul.io/http2curl"
)

// RequestLog ...
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		_procRequestLog(c)
		c.Next()
	}
}

func _procRequestLog(c *gin.Context) {
	txnID := l.LogTraceID()
	c.Set("x-request-id", txnID)                 // set to request param
	c.Writer.Header().Set("X-Request-Id", txnID) // set to request header

	curl, _ := http2curl.GetCurlCommand(c.Request)
	l.Logger(txnID, "").Infoln(fmt.Sprint(curl))
}

//Auth middleware
// func Auth(next gin.HandlerFunc) gin.HandlerFunc {
// 	return func(c gin.Context) error {
// 		appID := c.FormValue("appId")
// 		accessKey := c.FormValue("secretKey")
// 		fmt.Println("id,key", appID, accessKey)
// 		if appID == "id" && accessKey == "key" {
// 			fmt.Println("authenticated from middleware")
// 			return next(c)
// 		}
// 		return c.JSON(http.StatusBadRequest, "bad request")
// 	}
// }

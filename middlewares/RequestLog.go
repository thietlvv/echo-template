package middleware

import (
	l "billing/utils/logger"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"moul.io/http2curl"
)

// RequestLog ...
func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_procRequestLog(c)
		return next(c)
	}
}

func _procRequestLog(c echo.Context) {
	txnID := l.LogTraceID()
	c.Set("x-request-id", txnID)                  // set to request param
	c.Request().Header.Set("X-Request-Id", txnID) // set to request header

	curl, _ := http2curl.GetCurlCommand(c.Request())
	l.Logger(txnID, "").Infoln(fmt.Sprint(curl))
}

//Auth middleware
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		appID := c.FormValue("appId")
		accessKey := c.FormValue("secretKey")
		fmt.Println("id,key", appID, accessKey)
		if appID == "id" && accessKey == "key" {
			fmt.Println("authenticated from middleware")
			return next(c)
		}
		return c.JSON(http.StatusBadRequest, "bad request")
	}
}

package urls

import (
	middlewares "billing/middlewares"
	"billing/serv/monitor"
	"billing/serv/tpbank"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitUrlsBilling(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	// Init log global
	e.Use(middlewares.RequestLog)

	// Init group /billing
	gBilling := e.Group("/billing")

	// Init group /billing/m
	monitor.Routes(gBilling)

	// Init group /billing/tpbank
	tpbank.Routes(gBilling)
}

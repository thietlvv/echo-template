package monitor

import (
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	gMonitor := g.Group("/m")

	gMonitor.GET("/health_check", HealthCheck)
	gMonitor.GET("/load_test", LoadTest)
}

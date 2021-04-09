package tpbank

import (
	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) {
	gTpbank := g.Group("/tpbank")

	gTpbank.GET("", sayHello)
}

package tpbank

import (
	"billing/utils/response"

	"github.com/labstack/echo/v4"
)

func sayHello(c echo.Context) error {
	// ctx := c.Request().Context()
	// err := u.crawlerUseCase.CreateCoin(ctx)
	// if err != nil {
	// 	log.Error(err)
	// 	return response.Error(c, err)
	// }
	return response.Success(c, echo.Map{"data": "ok"})
}
